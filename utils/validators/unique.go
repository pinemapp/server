package validators

import (
	"fmt"
	"regexp"

	"github.com/pinem/server/db"

	"gopkg.in/go-playground/validator.v9"
)

var uniqueParamRegexp = regexp.MustCompile(`(.*)\((.*)\)`)

func uniqueValidator(fl validator.FieldLevel) bool {
	param := fl.Param()
	if param == "" {
		return false
	}
	data := uniqueParamRegexp.FindStringSubmatch(param)
	if len(data) != 3 {
		return false
	}
	tableName := data[1]
	colName := data[2]

	val := fl.Field().String()
	var count int
	err := db.ORM.Table(tableName).Where(fmt.Sprintf("%s = ?", colName), val).Count(&count).Error
	if err != nil {
		return false
	}

	return count == 0
}
