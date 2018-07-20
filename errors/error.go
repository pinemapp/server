package errors

import (
	gerrors "errors"
	"strings"

	"github.com/jinzhu/gorm"
)

const (
	ForeignKeyErrMsg = "violates foreign key constraint"
)

var (
	ErrInternalServer     = gerrors.New("errors_internal_server")
	ErrInvalidCredentials = gerrors.New("errors_invalid_credentials")
	ErrForbidden          = gerrors.New("errors_forbidden")
	ErrRecordNotFound     = gerrors.New("errors_not_found")
	ErrNotUnique          = gerrors.New("errors_already_exist")
	ErrForeignKey         = gerrors.New("errors_foreign_key")
)

func GetDBError(err error) error {
	if strings.Contains(err.Error(), ForeignKeyErrMsg) {
		return ErrForeignKey
	} else if gorm.IsRecordNotFoundError(err) {
		return ErrRecordNotFound
	}
	return err
}
