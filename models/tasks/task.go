package tasks

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/pinem/server/db"
	"github.com/pinem/server/errors"
	"github.com/pinem/server/models"
	"github.com/pinem/server/utils"
	"github.com/pinem/server/utils/messages"
	"github.com/pinem/server/utils/validators"
	"github.com/pinem/server/utils/validators/tasks"
)

func Create(c *gin.Context, msg *messages.Messages) (*models.Task, error) {
	var f taskvalidator.TaskForm
	c.Bind(&f)
	if err := validators.Validate(&f, msg); err != nil {
		return nil, err
	}

	boardID := utils.GetIntParam("board_id", c)
	if !isListExist(f.ListID, boardID) {
		return nil, errors.ErrRecordNotFound
	}

	var task models.Task
	lastTask := getLastTask(f.ListID, c)
	validators.Bind(&task, &f)
	task.BoardID = boardID
	task.Order = lastTask.Order + 1

	if err := db.ORM.Create(&task).Error; err != nil {
		return nil, errors.GetDBError(err)
	}
	return &task, nil
}

func Update(c *gin.Context, msg *messages.Messages) (*models.Task, error) {
	var f taskvalidator.UpdateTaskForm
	c.Bind(&f)
	if err := validators.Validate(&f, msg); err != nil {
		return nil, err
	}

	task, err := GetOneInBoard(c)
	if err != nil {
		return nil, errors.ErrRecordNotFound
	}

	var newTask models.Task
	validators.Bind(&newTask, &f)

	isMoveToNewList := newTask.ListID != 0 && task.ListID != newTask.ListID
	if isMoveToNewList && !isListExist(newTask.ListID, task.BoardID) {
		return nil, errors.ErrRecordNotFound
	}

	lastTask := getLastTask(newTask.ListID, c)
	if isMoveToNewList && newTask.Order == 0 {
		if lastTask.Order != 0 {
			newTask.Order = lastTask.Order + 1
		} else {
			newTask.Order = 1
		}
		f.Order = &newTask.Order
	}

	if isMoveToNewList {
		if lastTask.Order != 0 && newTask.Order > lastTask.Order+1 {
			return nil, errors.ErrOrderOutOfRange
		}
		if lastTask.Order == 0 && newTask.Order != 1 {
			return nil, errors.ErrOrderOutOfRange
		}
	} else if newTask.Order > lastTask.Order {
		return nil, errors.ErrOrderOutOfRange
	}

	err = update(&f, &newTask, task, c, isMoveToNewList)
	if err != nil {
		return nil, err
	}

	return task, nil
}

func Delete(c *gin.Context) error {
	task, err := GetOneInBoard(c)
	if err != nil {
		return errors.ErrRecordNotFound
	}

	err = db.Transaction(db.ORM, func(tx *gorm.DB) error {
		err := tx.Delete(task).Error
		if err != nil {
			return errors.GetDBError(err)
		}

		err = reorder(tx, task.BoardID, task.ListID, task.Order, 100000, 1)
		if err != nil {
			return errors.GetDBError(err)
		}

		return nil
	})

	if err != nil {
		return err
	}
	return nil
}

func getLastTask(listID uint, c *gin.Context) models.Task {
	var lastTask models.Task
	Scope(c).Where("tasks.list_id = ?", listID).
		Order("tasks.order DESC").First(&lastTask)
	return lastTask
}

func update(f *taskvalidator.UpdateTaskForm, newTask, task *models.Task, c *gin.Context, isMoveToNewList bool) error {
	err := db.Transaction(db.ORM, func(tx *gorm.DB) error {
		neededUpdate := false
		newOrder := newTask.Order
		oldOrder := task.Order
		newListID := newTask.ListID
		oldListID := task.ListID
		validators.Bind(task, f)

		if (newOrder != 0 && oldOrder != newOrder) ||
			(newListID != 0 && oldListID != newListID) {
			neededUpdate = true
		}

		if neededUpdate {
			boardID := utils.GetIntParam("board_id", c)
			if isMoveToNewList {
				// re-order task in old list
				if err := reorder(tx, boardID, oldListID, oldOrder, 100000, 1); err != nil {
					return err
				}
				// re-order task in current list
				if err := reorder(tx, boardID, task.ListID, newOrder-1, 100000, -1); err != nil {
					return err
				}
			} else {
				// re-order task in current list
				coe := -1
				if oldOrder < newOrder {
					coe = 1
				}
				min, max := utils.GetOrderRange(oldOrder, newOrder)
				err := reorder(tx, boardID, task.ListID, min, max, coe)
				if err != nil {
					return err
				}
			}
		}

		if err := tx.Save(task).Error; err != nil {
			return errors.GetDBError(err)
		}
		return nil
	})
	return err
}

func isListExist(listID, boardID uint) bool {
	var count int
	err := db.ORM.Table("lists").Joins("JOIN boards ON boards.id = lists.board_id").
		Where("boards.id = ? AND lists.id = ?", boardID, listID).Count(&count).Error
	if err != nil {
		return false
	}
	return count > 0
}

func reorder(tx *gorm.DB, boardID, listID uint, min, max, coe int) error {
	sql := "tasks.board_id = ? AND tasks.list_id = ? AND tasks.order > ? AND tasks.order < ?"
	err := tx.Model(models.Task{}).Where(sql, boardID, listID, min, max).
		UpdateColumn("order", gorm.Expr("tasks.order - ?", coe)).Error
	if err != nil {
		return errors.GetDBError(err)
	}
	return nil
}
