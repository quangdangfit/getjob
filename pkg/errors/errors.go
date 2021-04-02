package errors

import (
	"fmt"

	"github.com/pkg/errors"
)

//Reference article:
//https://hackernoon.com/golang-handling-errors-gracefully-8e27f1db729f

type CustomError struct {
	code         ErrorCode
	wrappedError error
	context      errorContext
}

func (err CustomError) Error() string {
	return err.wrappedError.Error()
}

func (err CustomError) Stacktrace() string {
	return fmt.Sprintf("%+v\n", err.wrappedError)
}

func (err CustomError) GetErrorCode() ErrorCode {
	return err.code
}

// New creates a no type error
func New(errCode ErrorCode, msg string) error {
	return CustomError{code: errCode, wrappedError: errors.New(msg)}
}

// Newf creates a no type error with formatted message
func Newf(errCode ErrorCode, msg string, args ...interface{}) error {
	return CustomError{code: errCode, wrappedError: errors.New(fmt.Sprintf(msg, args...))}
}

// Wrap wrans an error with a string
func Wrap(err error, msg string) error {
	return Wrapf(err, msg)
}

// Cause gives the original error
func Cause(err error) error {
	return errors.Cause(err)
}

// Wrapf wraps an error with format string
func Wrapf(err error, msg string, args ...interface{}) error {
	wrappedError := errors.Wrapf(err, msg, args...)
	if customErr, ok := err.(CustomError); ok {
		return CustomError{
			code:         customErr.code,
			wrappedError: wrappedError,
			context:      customErr.context,
		}
	}

	return CustomError{code: ECUnknown, wrappedError: wrappedError}
}

// Get Stacktrace of error
func Stack(err error) string {
	if customErr, ok := err.(CustomError); ok {
		return fmt.Sprintf("%+v\n", customErr.wrappedError)
	}
	return fmt.Sprintf("%+v\n", errors.WithStack(err))
}
