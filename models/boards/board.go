package boards

import (
	"github.com/gin-gonic/gin"
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

	user := auth.GetUserFromContext(c)
	board := models.Board{
		Name:   f.Name,
		Desc:   f.Desc,
		UserID: user.ID,
		Public: f.Public,
	}
	if err := db.ORM.Create(&board).Error; err != nil {
		return nil, errors.ErrInternalServer
	}
	return &board, nil
}

func Update(c *gin.Context, msg *messages.Messages) (*models.Board, error) {
	board, err := GetOneForUser(c)
	if err != nil {
		return nil, err
	}

	var f boardvalidator.BoardForm
	c.Bind(&f)
	if err := validators.Validate(&f, msg); err != nil {
		return nil, err
	}

	board.Desc = f.Desc
	board.Name = f.Name
	board.Public = f.Public

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

	if err := db.ORM.Delete(&board).Error; err != nil {
		return err
	}
	return nil
}
