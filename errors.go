package errors

import (
	"errors"
	"fmt"
)

// As calls errors.As from the standard library.
func As(err error, target interface{}) bool {
	return errors.As(err, target)
}

// Is calls errors.Is from the standard library.
func Is(err, target error) bool {
	return errors.Is(err, target)
}

// New calls errors.New from the standard library.
func New(text string) error {
	return errors.New(text)
}

// Unwrap calls errors.Unwrap from the standard library.
func Unwrap(err error) error {
	return errors.Unwrap(err)
}

// Wrap returns the provided error annotated with the supplied message using
// fmt.Errorf. If err is nil, Wrap returns nil.
func Wrap(err error, message string) error {
	if err == nil {
		return nil
	}
	return fmt.Errorf("%v: %w", message, err)
}

// Wrapf returns the provided error annotated with the supplied message using
// fmt.Errorf. If err is nil, Wrapf returns nil.
func Wrapf(err error, format string, args ...interface{}) error {
	return Wrap(err, fmt.Sprintf(format, args...))
}
