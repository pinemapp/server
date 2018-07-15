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

	color := f.Color
	boardID := utils.GetIntParam("board_id", c)
	if color == "" {
		color = models.DefaultListColor
	}
	list := models.List{
		Name:    f.Name,
		Color:   color,
		BoardID: boardID,
		Order:   lastList.Order + 1,
	}
	err := db.ORM.Create(&list).Error
	if err != nil {
		return nil, errors.ErrNotFound
	}
	return &list, nil
}

func Update(c *gin.Context, msg *messages.Messages) (*models.List, error) {
	var f listvalidator.UpdateListForm
	c.Bind(&f)
	if err := validators.Validate(&f, msg); err != nil {
		return nil, err
	}

	list, err := GetOneInBoard(c)
	if err != nil {
		return nil, errors.ErrNotFound
	}

	var lastList models.List
	Scope(c).Order("lists.order DESC").First(&lastList)
	if f.Order > lastList.Order {
		msg.ErrorT("order", errors.ErrOrderOutOfRange)
		return nil, errors.ErrOrderOutOfRange
	}

	err = update(&f, list, c)
	if err != nil {
		return nil, err
	}
	return list, nil
}

func Delete(c *gin.Context) error {
	list, err := GetOneInBoard(c)
	if err != nil {
		return errors.ErrNotFound
	}

	err = db.Transaction(db.ORM, func(tx *gorm.DB) error {
		err := tx.Delete(list).Error
		if err != nil {
			return errors.ErrInternalServer
		}

		// TODO: find a way to get last order
		err = reorder(tx, list.BoardID, list.Order, 100000, 1)
		if err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return err
	}
	return nil
}

func update(f *listvalidator.UpdateListForm, list *models.List, c *gin.Context) error {
	err := db.Transaction(db.ORM, func(tx *gorm.DB) error {
		neededUpdate := false
		newOrder := f.Order
		oldOrder := list.Order
		list.Name = f.Name

		newList := models.List{
			Name:  f.Name,
			Color: f.Color,
		}
		if newOrder != 0 {
			newList.Order = newOrder
		}
		if newOrder != 0 && oldOrder != newOrder {
			neededUpdate = true
		}

		if neededUpdate {
			coe := -1
			if oldOrder < newOrder {
				coe = 1
			}
			boardID := utils.GetIntParam("board_id", c)
			min, max := utils.GetOrderRange(oldOrder, newOrder)
			err := reorder(tx, boardID, min, max, coe)
			if err != nil {
				return err
			}
		}

		err := tx.Model(list).Updates(newList).Error
		if err != nil {
			return errors.ErrNotFound
		}
		return nil
	})

	return err
}

func reorder(tx *gorm.DB, boardID uint, min, max, coe int) error {
	sql := "board_id = ? AND lists.order > ? AND lists.order < ?"
	err := tx.Model(models.List{}).Where(sql, boardID, min, max).
		UpdateColumn("order", gorm.Expr("lists.order - ?", coe)).Error
	if err != nil {
		return errors.ErrInternalServer
	}
	return nil
}
