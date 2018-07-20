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

	lastTask := getLastTask(f.ListID, c)
	task := models.Task{
		BoardID: boardID,
		ListID:  f.ListID,
		Name:    f.Name,
		Desc:    f.Desc,
		StartAt: f.StartAt,
		EndAt:   f.EndAt,
		Order:   lastTask.Order + 1,
	}
	if err := db.ORM.Create(&task).Error; err != nil {
		return nil, errors.ErrRecordNotFound
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

	isMoveToNewList := f.ListID != 0 && task.ListID != f.ListID
	if isMoveToNewList && !isListExist(f.ListID, task.BoardID) {
		return nil, errors.ErrRecordNotFound
	}

	if isMoveToNewList && f.Order == 0 {
		f.Order = 1
	}

	lastTask := getLastTask(f.ListID, c)
	if isMoveToNewList {
		if lastTask.Order != 0 && f.Order > lastTask.Order {
			return nil, errors.ErrOrderOutOfRange
		}
		if lastTask.Order == 0 && f.Order != 1 {
			return nil, errors.ErrOrderOutOfRange
		}
	} else if f.Order > lastTask.Order {
		return nil, errors.ErrOrderOutOfRange
	}

	err = update(&f, task, c, isMoveToNewList)
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
			return errors.ErrInternalServer
		}

		err = reorder(tx, task.BoardID, task.ListID, task.Order, 100000, 1)
		if err != nil {
			return errors.ErrInternalServer
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

func update(f *taskvalidator.UpdateTaskForm, task *models.Task, c *gin.Context, isMoveToNewList bool) error {
	err := db.Transaction(db.ORM, func(tx *gorm.DB) error {
		listID := f.ListID
		if listID == 0 {
			listID = task.ListID
		}

		neededUpdate := false
		newOrder := f.Order
		oldOrder := task.Order

		newListID := f.ListID
		oldListID := task.ListID

		newTask := models.Task{
			Name:    f.Name,
			ListID:  listID,
			Desc:    f.Desc,
			StartAt: f.StartAt,
			EndAt:   f.EndAt,
		}

		if newOrder != 0 {
			newTask.Order = newOrder
		}
		if (newOrder != 0 && oldOrder != newOrder) ||
			(newListID != 0 && oldListID != newListID) {
			neededUpdate = true
		}

		if neededUpdate {
			boardID := utils.GetIntParam("board_id", c)
			if isMoveToNewList {
				// re-order task in old list
				err := reorder(tx, boardID, task.ListID, task.Order, 100000, 1)
				if err != nil {
					return err
				}
			}

			// re-order task in current list
			coe := -1
			if oldOrder < newOrder {
				coe = 1
			}
			min, max := utils.GetOrderRange(oldOrder, newOrder)
			err := reorder(tx, boardID, newTask.ListID, min, max, coe)
			if err != nil {
				return err
			}
		}

		err := tx.Model(task).Updates(newTask).Error
		if err != nil {
			return errors.ErrRecordNotFound
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
		return errors.ErrInternalServer
	}
	return nil
}
