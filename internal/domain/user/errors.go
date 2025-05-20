package user

import "errors"

var (
	ErrInsufficientBalance = errors.New("user: insufficient balance to transfer")
	ErrUserNotAuthorized   = errors.New("user: user not authorized to perform this action")
)
