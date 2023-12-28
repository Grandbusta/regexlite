package regexlite

import (
	"errors"
	"net/mail"
	"net/url"
	"strings"
)

func ValidateString(s string) *sWrapper {
	return &sWrapper{value: s}
}

func (w *sWrapper) Min(n int) *sWrapper {
	size := len(w.value)
	if size < n {
		if w.err == nil {
			w.err = ErrTooSmall
		} else {
			w.err = errors.Join(w.err, ErrTooSmall)
		}
	}
	return w
}

func (w *sWrapper) Max(n int) *sWrapper {
	size := len(w.value)
	if size > n {
		if w.err == nil {
			w.err = ErrTooBig
		} else {
			w.err = errors.Join(w.err, ErrTooBig)
		}
	}
	return w
}

func (w *sWrapper) IsEmail() *sWrapper {
	email, err := mail.ParseAddress(w.value) // Try to use what std gives, avoid regexes if possible.
	if err != nil || strings.ToLower(w.value) != email.String() {
		if w.err == nil {
			w.err = ErrNotEmail
		} else {
			w.err = errors.Join(w.err, ErrNotEmail)
		}
	}
	return w
}

func (w *sWrapper) IsURL() *sWrapper {
	urlData, err := url.Parse(w.value)
	if err != nil || w.value != urlData.String() {
		if w.err == nil {
			w.err = ErrNotURL
		} else {
			w.err = errors.Join(w.err, ErrNotURL)
		}
	}
	return w
}

// Returns an error with explanation if checked value did not pass requirements.
func (w *sWrapper) Check() error {
	return w.err
}
