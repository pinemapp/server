package listvalidator

type ListForm struct {
	Name  string `json:"name" validate:"required,max=150"`
	Color string `json:"color" validate:"omitempty,hexcolor"`
}

type UpdateListForm struct {
	Name  string `json:"name" validate:"omitempty,max=150"`
	Color string `json:"color" validate:"omitempty,hexcolor"`
	Order int    `json:"order" validate:"omitempty,numeric"`
}
