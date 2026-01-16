package service

import "errors"

var (
	ErrNovelaNotFound = errors.New("novela not found")
	ErrInternal       = errors.New("internal server error")
)