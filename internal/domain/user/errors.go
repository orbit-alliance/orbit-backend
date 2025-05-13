package user

import "errors"

var (
	ErrInsufficientBalance   = errors.New("user: insufficient balance to transfer")
	ErrFailToCreate42Token   = errors.New("user: fail to create 42 token")
	ErrFailToGetLocationIn42 = errors.New("user: fail to get location in 42")
	ErrFailToGetUserIn42     = errors.New("user: fail to get user in 42")
)
