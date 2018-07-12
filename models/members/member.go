package members

import (
	"github.com/gin-gonic/gin"
	"github.com/pinem/server/db"
	"github.com/pinem/server/errors"
	"github.com/pinem/server/models"
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

	ok, err := isMember(f.UserID, c)
	if err != nil {
		return nil, errors.ErrNotFound
	} else if ok {
		return nil, errors.ErrMemberAlreadyInBoard
	}

	boardID := utils.GetIntParam("board_id", c)
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

func Update(c *gin.Context, msg *messages.Messages) (*models.BoardUser, error) {
	var f membervalidator.UpdateMemberForm
	c.Bind(&f)

	if err := validators.Validate(&f, msg); err != nil {
		return nil, err
	}

	member, err := GetOneInBoard(c)
	if err != nil {
		return nil, err
	}

	member.Role = f.Role
	if err := db.ORM.Save(member).Error; err != nil {
		return nil, errors.ErrInternalServer
	}

	return member, nil
}

func Delete(c *gin.Context) error {
	member, err := GetOneInBoard(c)
	if err != nil {
		return err
	}

	if err := db.ORM.Delete(member).Error; err != nil {
		return errors.ErrInternalServer
	}
	return nil
}

func isMember(userID uint, c *gin.Context) (bool, error) {
	var count int
	if err := Scope(c).Model(&models.BoardUser{}).Where("board_users.user_id = ?", userID).Count(&count).Error; err != nil {
		return false, err
	}
	return count > 0, nil
}
