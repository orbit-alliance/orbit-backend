package services

import "errors"

var (
	ErrGoodActionNotFound  = errors.New("good action not found")
	ErrUserNotHasStartDate = errors.New("user does not have a start date")
)
