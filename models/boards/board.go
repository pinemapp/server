package boards

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/pinem/server/db"
	"github.com/pinem/server/errors"
	"github.com/pinem/server/models"
	"github.com/pinem/server/utils/auth"
	"github.com/pinem/server/utils/messages"
	"github.com/pinem/server/utils/validators"
	"github.com/pinem/server/utils/validators/boards"
)

func Create(c *gin.Context, msg *messages.Messages) (*models.Board, error) {
	var f boardvalidator.BoardForm
	c.Bind(&f)
	if err := validators.Validate(&f, msg); err != nil {
		return nil, err
	}

	board, err := create(&f, c)
	if err != nil {
		return nil, errors.ErrInternalServer
	}
	return board, nil
}

func Update(c *gin.Context, msg *messages.Messages) (*models.Board, error) {
	board, err := GetOneForUser(c)
	if err != nil {
		return nil, err
	}

	var f boardvalidator.UpdateBoardForm
	c.Bind(&f)
	if err := validators.Validate(&f, msg); err != nil {
		return nil, err
	}

	validators.Bind(board, &f)
	if err := db.ORM.Save(board).Error; err != nil {
		return nil, errors.ErrInternalServer
	}
	return board, nil
}

func Delete(c *gin.Context) error {
	board, err := GetOneForUser(c)
	if err != nil {
		return err
	}

	err = db.Transaction(db.ORM, func(tx *gorm.DB) error {
		if err := tx.Delete(models.Task{}, "board_id = ?", board.ID).Error; err != nil {
			return err
		}
		if err := tx.Delete(models.List{}, "board_id = ?", board.ID).Error; err != nil {
			return err
		}
		if err := tx.Delete(models.BoardUser{}, "board_id = ?", board.ID).Error; err != nil {
			return err
		}
		if err := tx.Delete(&board).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

func create(f *boardvalidator.BoardForm, c *gin.Context) (*models.Board, error) {
	user := auth.GetUserFromContext(c)
	var board models.Board
	validators.Bind(&board, f)

	err := db.Transaction(db.ORM, func(tx *gorm.DB) error {
		if err := tx.Create(&board).Error; err != nil {
			return errors.ErrInternalServer
		}
		boardUser := models.BoardUser{
			UserID:  user.ID,
			BoardID: board.ID,
			Role:    models.BoardOwner,
		}
		if err := tx.Create(&boardUser).Error; err != nil {
			return errors.ErrInternalServer
		}
		board.Members = append(board.Members, boardUser)
		return nil
	})

	if err != nil {
		return nil, err
	}
	return &board, nil
}
