package members

import (
	"github.com/gin-gonic/gin"
	"github.com/pinem/server/db"
	"github.com/pinem/server/errors"
	"github.com/pinem/server/models"
	"github.com/pinem/server/models/boards"
	"github.com/pinem/server/utils"
	"github.com/pinem/server/utils/messages"
	"github.com/pinem/server/utils/validators"
	"github.com/pinem/server/utils/validators/members"
)

func Add(c *gin.Context, msg *messages.Messages) (*models.BoardUser, error) {
	var f membervalidator.MemberForm
	c.Bind(&f)

	if err := validators.Validate(&f, msg); err != nil {
		return nil, err
	}

	boardID := utils.GetIntParam("board_id", c)
	ok, err := isMember(boardID, f.UserID, c)
	if err != nil {
		return nil, errors.ErrNotFound
	} else if ok {
		return nil, errors.ErrMemberAlreadyInBoard
	}

	member := models.BoardUser{
		UserID:  f.UserID,
		BoardID: boardID,
		Role:    f.Role,
	}
	if err := db.ORM.Create(&member).Error; err != nil {
		return nil, errors.ErrUserNotExist
	}

	return &member, nil
}

func isMember(boardID uint, userID uint, c *gin.Context) (bool, error) {
	var count int
	if err := boards.Scope(c).Table("boards").Where("boards.id = ? AND board_users.user_id = ?", boardID, userID).Count(&count).Error; err != nil {
		return false, err
	}
	return count > 0, nil
}
