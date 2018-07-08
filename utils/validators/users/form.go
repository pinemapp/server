package uservalidator

type RegisterForm struct {
	Name        string `json:"name" validate:"required,max=150,unique=users(name)"`
	Email       string `json:"email" validate:"required,email,max=150,unique=users(email)"`
	Password    string `json:"password" validate:"required,min=8,max=150"`
	RedirectUrl string `json:"redirect_url" validate:"required"`
}
