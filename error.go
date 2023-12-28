package regexlite

import "errors"

var (
	ErrTooBig   = errors.New("regexlite: provided value is too big")
	ErrTooSmall = errors.New("regexlite: provided value is too small")
)
