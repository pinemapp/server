package projectvalidator

type ProjectForm struct {
	Name   string  `json:"name" validate:"required,max=150"`
	Desc   *string `json:"desc,omitempty" validate:"omitempty,max=1000"`
	Public bool    `json:"public,omitempty" validate:"-"`
	TeamID *uint   `json:"team_id,omitempty" validate:"-"`
}

type UpdateProjectForm struct {
	Name   *string `json:"name,omitempty" validate:"omitempty,max=150"`
	Desc   *string `json:"desc,omitempty" validate:"omitempty,max=1000"`
	Public *bool   `json:"public,omitempty" validate:"-"`
	TeamID *uint   `json:"team_id,omitempty" validate:"-"`
}
