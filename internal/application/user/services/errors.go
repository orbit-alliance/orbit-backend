package services

import "errors"

var (
	ErrGoodActionNotFound  = errors.New("Good action not found")
	ErrUserNotHasStartDate = errors.New("User does not have a start date")
)
