package errors

import gerrors "errors"

var (
	ErrUserNotExist         = gerrors.New("errors_user_not_exist")
	ErrMemberAlreadyInBoard = gerrors.New("errors_member_already_in_board")
)
