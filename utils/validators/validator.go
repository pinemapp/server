package validators

import (
	"fmt"
	"reflect"
	"time"

	"github.com/pinem/server/utils/messages"
	validator "gopkg.in/go-playground/validator.v9"
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
			if err.Tag() == "unique" {
				msg.AddErrorT(err.Field(), fmt.Sprintf("errors_%s", err.Tag()))
			} else if err.Param() != "" {
				msg.AddErrorTf(err.Field(), fmt.Sprintf("errors_%s", err.Tag()), err.Param())
			} else {
				msg.AddErrorT(err.Field(), fmt.Sprintf("errors_%s", err.Tag()))
			}
		}
		return err
	}
	return nil
}

func Bind(model interface{}, f interface{}) {
	modelIndirect := reflect.Indirect(reflect.ValueOf(model))
	formElem := reflect.ValueOf(f).Elem()
	tForm := formElem.Type()
	for i := 0; i < formElem.NumField(); i++ {
		tag := tForm.Field(i).Tag
		if tag.Get("omit") != "true" {
			modelField := modelIndirect.FieldByName(tForm.Field(i).Name)
			if modelField.IsValid() {
				formField := formElem.Field(i)
				if formField.IsValid() {
					if isDirty(modelField, formField) {
						modelField.Set(formField)
					}
				}
			}
		}
	}
}

func isDirty(f, ff reflect.Value) bool {
	switch f.Kind() {
	case reflect.String:
		return f.String() != ff.String()
	case reflect.Bool:
		return f.Bool() != ff.Bool()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return f.Int() != ff.Int()
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return f.Uint() != ff.Uint()
	case reflect.Float32, reflect.Float64:
		return f.Float() != f.Float()
	case reflect.Interface, reflect.Ptr:
		if ff.IsNil() {
			return false
		}
		if f.IsNil() && !ff.IsNil() {
			return true
		}
		return isDirty(f.Elem(), ff.Elem())
	case reflect.Struct:
		if t, ok := f.Interface().(time.Time); ok {
			tt := ff.Interface().(time.Time)
			return !t.Equal(tt)
		}
	}
	return false
}
