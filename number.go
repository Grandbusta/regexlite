package regexlite

import "errors"

func ValidateNumber[T number](n T) *nWrapper[T] {
	return &nWrapper[T]{value: n}
}

func (w *nWrapper[T]) Min(n T) *nWrapper[T] {
	if w.value < n {
		if w.err == nil {
			w.err = ErrTooSmall
		} else {
			w.err = errors.Join(w.err, ErrTooSmall)
		}
	}
	return w
}

func (w *nWrapper[T]) Max(n T) *nWrapper[T] {
	if w.value > n {
		if w.err == nil {
			w.err = ErrTooBig
		} else {
			w.err = errors.Join(w.err, ErrTooBig)
		}
	}
	return w
}

// Returns an error with explanation if checked value did not pass requirements.
func (w *nWrapper[T]) Check() error {
	return w.err
}
