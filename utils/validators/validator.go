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
						mKind, fKind := modelField.Kind(), formField.Kind()
						if mKind == reflect.Ptr && fKind != reflect.Ptr {
							modelField.Set(formField.Addr())
						} else if mKind == fKind {
							modelField.Set(formField)
						} else if !formField.IsNil() {
							modelField.Set(formField.Elem())
						}
					}
				}
			}
		}
	}
}

func isDirty(f, ff reflect.Value) bool {
	fKind, ffKind, fff := f.Kind(), ff.Kind(), ff
	if fKind != ffKind {
		if fKind != reflect.Ptr && fKind != reflect.Interface && fKind != reflect.Struct {
			if ff.IsNil() {
				return false
			}
			fff = ff.Elem()
		}
	}

	switch fKind {
	case reflect.String:
		return f.String() != fff.String()
	case reflect.Bool:
		return f.Bool() != fff.Bool()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return f.Int() != fff.Int()
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return f.Uint() != fff.Uint()
	case reflect.Float32, reflect.Float64:
		return f.Float() != fff.Float()
	case reflect.Interface, reflect.Ptr:
		if ffKind == reflect.Ptr || ffKind == reflect.Interface {
			if ff.IsNil() {
				return false
			}
			if f.IsNil() {
				return !ff.IsNil()
			}
			return isDirty(f.Elem(), ff.Elem())
		}
		if f.IsNil() {
			return ff.Interface() != nil
		}
		return isDirty(f.Elem(), ff)
	case reflect.Struct:
		if t, ok := f.Interface().(time.Time); ok {
			tt := ff.Interface().(time.Time)
			return !t.Equal(tt)
		}
	}
	return false
}
