package membervalidator

import "github.com/pinem/server/models"

type MemberForm struct {
	UserID uint               `json:"user_id" validate:"required"`
	Role   models.ProjectRole `json:"role" validate:"required,oneof=owner member visitor"`
}

type UpdateMemberForm struct {
	Role models.ProjectRole `json:"role" validate:"required,oneof=owner member visitor"`
}
