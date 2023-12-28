package regexlite

import "errors"

var (
	ErrTooBig   = errors.New("regexlite: provided value is too big")
	ErrTooSmall = errors.New("regexlite: provided value is too small")
	ErrNotEmail = errors.New("regexlite: provided value is not a valid address")
	ErrNotURL   = errors.New("regexlite: provided value is not a valid url")
)
