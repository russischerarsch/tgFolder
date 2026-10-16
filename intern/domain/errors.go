package domain

import "errors"

var (
	ErrTooLongName = errors.New("the name you applied is too long")
)
