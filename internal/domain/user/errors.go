package user

import "errors"

var (
	ErrInsufficientBalance = errors.New("user: insufficient balance to transfer")
)
