package teamvalidator

type TeamForm struct {
	Name    string  `json:"name" validate:"required,max=150"`
	Slug    *string `json:"slug,omitempty"`
	Desc    *string `json:"desc,omitempty" validate:"omitempty,max=1000"`
	Website *string `json:"website,omitempty" validate:"omitempty,url"`
}

type UpdateTeamForm struct {
	Name    *string `json:"name,omitempty" validate:"omitempty,max=150"`
	Slug    *string `json:"slug,omitempty"`
	Desc    *string `json:"desc,omitempty" validate:"omitempty,max=1000"`
	Website *string `json:"website,omitempty" validate:"omitempty,url"`
}
