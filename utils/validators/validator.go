package validators

import (
	"fmt"

	"github.com/pinem/server/utils/messages"
	"gopkg.in/go-playground/validator.v9"
)

var validate *validator.Validate

func init() {
	validate = validator.New()
	validate.RegisterValidation("unique", uniqueValidator)
}

func Validate(f interface{}, msg *messages.Messages) error {
	err := validate.Struct(f)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			msg.AddErrorT(err.Field(), fmt.Sprintf("errors_%s", err.Tag()))
		}
		return err
	}
	return nil
}
