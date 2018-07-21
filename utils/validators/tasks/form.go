package taskvalidator

import "time"

type TaskForm struct {
	Name    string     `json:"name" validate:"required,max=150"`
	ListID  uint       `json:"list_id" validate:"required"`
	Desc    *string    `json:"desc,omitempty" validate:"omitempty,max=1000"`
	StartAt *time.Time `json:"start_at,omitempty" validate:"omitempty,ltefield=EndAt"`
	EndAt   *time.Time `json:"end_at,omitempty" validate:"omitempty,gtefield=StartAt"`
}

type UpdateTaskForm struct {
	Name    *string    `json:"name,omitempty" validate:"omitempty,max=150"`
	Desc    *string    `json:"desc,omitempty" validate:"omitempty,max=1000"`
	ListID  *uint      `json:"list_id,omitempty" validate:"omitempty"`
	Order   *int       `json:"order,omitempty" validate:"omitempty,numeric"`
	StartAt *time.Time `json:"start_at,omitempty" validate:"omitempty,ltefield=EndAt"`
	EndAt   *time.Time `json:"end_at,omitempty" validate:"omitempty,gtefield=StartAt"`
}
