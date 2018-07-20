package listvalidator

type ListForm struct {
	Name  string  `json:"name" validate:"required,max=150"`
	Color *string `json:"color,omitempty" validate:"omitempty,hexcolor"`
}

type UpdateListForm struct {
	Name  *string `json:"name,omitempty" validate:"omitempty,max=150"`
	Color *string `json:"color,omitempty" validate:"omitempty,hexcolor"`
	Order *int    `json:"order,omitempty" validate:"omitempty,numeric"`
}
