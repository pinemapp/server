package boardvalidator

type BoardForm struct {
	Name   string `json:"name" validate:"required,max=150"`
	Desc   string `json:"desc" validate:"max=1000"`
	Public bool   `json:"public" validate:"-"`
}
