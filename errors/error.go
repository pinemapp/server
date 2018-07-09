package errors

import gerrors "errors"

var (
	ErrInternalServer     = gerrors.New("errors_internal_server")
	ErrInvalidCredentials = gerrors.New("errors_invalid_credentials")
	ErrForbidden          = gerrors.New("errors_forbidden")
)
