package user

import "errors"

var (
	ErrInsufficientBalance     = errors.New("user: insufficient balance to transfer")
	ErrUserNotAuthorized       = errors.New("user: user not authorized to perform this action")
	ErrSenderNotFound          = errors.New("user: sender not found")
	ErrReceiverNotFound        = errors.New("user: receiver not found")
	ErrUserNotFound            = errors.New("user: user not found")
	ErrUserAlreadyExists       = errors.New("user: user already exists")
	ErrWalletAlreadyRegistered = errors.New("user: wallet already registered")
	ErrSameWalletAddress       = errors.New("user: same wallet address provided")
	ErrInvalidEventType        = errors.New("user: invalid event type")
)
