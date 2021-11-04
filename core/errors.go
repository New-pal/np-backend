package core

import "errors"

var (
	errEmptyPassword  = errors.New("password is empty")
	ErrRecordNotFound = errors.New("record not found")
)
