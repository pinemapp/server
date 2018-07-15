package taskvalidator

import "time"

type TaskForm struct {
	Name    string     `json:"name" validate:"required,max=150"`
	Desc    string     `json:"desc" validate:"omitempty,max=1000"`
	ListID  uint       `json:"list_id" validate:"required"`
	StartAt *time.Time `json:"start_at" validate:"omitempty,ltefield=EndAt"`
	EndAt   *time.Time `json:"end_at" validate:"omitempty,gtefield=StartAt"`
}

type UpdateTaskForm struct {
	Name    string     `json:"name" validate:"omitempty,max=150"`
	Desc    string     `json:"desc" validate:"omitempty,max=1000"`
	ListID  uint       `json:"list_id" validate:"omitempty"`
	Order   int        `json:"order" validate:"omitempty,numeric"`
	StartAt *time.Time `json:"start_at" validate:"omitempty,ltefield=EndAt"`
	EndAt   *time.Time `json:"end_at" validate:"omitempty,gtefield=StartAt"`
}
