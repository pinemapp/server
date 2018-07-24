package lists

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/pinem/server/db"
	"github.com/pinem/server/errors"
	"github.com/pinem/server/models"
	"github.com/pinem/server/utils"
	"github.com/pinem/server/utils/messages"
	"github.com/pinem/server/utils/validators"
	"github.com/pinem/server/utils/validators/lists"
)

func Create(c *gin.Context, msg *messages.Messages) (*models.List, error) {
	var f listvalidator.ListForm
	c.Bind(&f)
	if err := validators.Validate(&f, msg); err != nil {
		return nil, err
	}

	var lastList models.List
	Scope(c).Order("lists.order DESC").First(&lastList)

	var list models.List
	validators.Bind(&list, &f)
	if list.Color == nil {
		list.Color = &models.DefaultListColor
	}
	list.Order = lastList.Order + 1
	list.ProjectID = utils.GetIntParam("project_id", c)

	err := db.ORM.Create(&list).Error
	if err != nil {
		return nil, errors.ErrRecordNotFound
	}
	return &list, nil
}

func Update(c *gin.Context, msg *messages.Messages) (*models.List, error) {
	var f listvalidator.UpdateListForm
	c.Bind(&f)
	if err := validators.Validate(&f, msg); err != nil {
		return nil, err
	}

	list, err := GetOneInProject(c)
	if err != nil {
		return nil, errors.ErrRecordNotFound
	}

	var lastList models.List
	Scope(c).Order("lists.order DESC").First(&lastList)
	if f.Order != nil {
		if *f.Order > lastList.Order {
			msg.ErrorT("order", errors.ErrOrderOutOfRange)
			return nil, errors.ErrOrderOutOfRange
		}
	}

	err = update(&f, list, c)
	if err != nil {
		return nil, err
	}
	return list, nil
}

func Delete(c *gin.Context) error {
	list, err := GetOneInProject(c)
	if err != nil {
		return errors.ErrRecordNotFound
	}

	err = db.Transaction(db.ORM, func(tx *gorm.DB) error {
		if err := tx.Delete(models.Task{}, "list_id = ?", list.ID).Error; err != nil {
			return err
		}
		if err := tx.Delete(list).Error; err != nil {
			return err
		}
		// TODO: find a way to get last order
		if err := reorder(tx, list.ProjectID, list.Order, 100000, 1); err != nil {
			return err
		}
		return nil
	})
	return errors.GetDBError(err)
}

func update(f *listvalidator.UpdateListForm, list *models.List, c *gin.Context) error {
	err := db.Transaction(db.ORM, func(tx *gorm.DB) error {
		oldOrder := list.Order
		validators.Bind(list, f)

		if list.Order != oldOrder {
			coe := -1
			if list.Order > oldOrder {
				coe = 1
			}
			projectID := utils.GetIntParam("project_id", c)
			min, max := utils.GetOrderRange(oldOrder, list.Order)
			err := reorder(tx, projectID, min, max, coe)
			if err != nil {
				return err
			}
		}

		if err := tx.Save(list).Error; err != nil {
			return err
		}
		return nil
	})
	return errors.GetDBError(err)
}

func reorder(tx *gorm.DB, projectID uint, min, max, coe int) error {
	sql := "project_id = ? AND lists.order > ? AND lists.order < ?"
	err := tx.Model(models.List{}).Where(sql, projectID, min, max).
		UpdateColumn("order", gorm.Expr("lists.order - ?", coe)).Error
	if err != nil {
		return err
	}
	return nil
}
