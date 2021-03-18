package errors

type errorContext struct {
	Field   string
	Message string
}

// AddErrorContext adds a context to an error
func AddErrorContext(err error, field, message string) error {
	context := errorContext{Field: field, Message: message}
	if customErr, ok := err.(CustomError); ok {
		return CustomError{
			code:         customErr.code,
			wrappedError: customErr.wrappedError,
			context:      context,
		}
	}

	return CustomError{
		code:         ECUnknown,
		wrappedError: err,
		context:      context,
	}
}

// GetErrorContext returns the error context
func GetErrorContext(err error) map[string]string {
	emptyContext := errorContext{}
	if customErr, ok := err.(CustomError); ok || customErr.context != emptyContext {

		return map[string]string{
			"field":   customErr.context.Field,
			"message": customErr.context.Message,
		}
	}

	return nil
}

// GetType returns the error type
func GetCode(err error) ErrorCode {
	if customErr, ok := err.(CustomError); ok {
		return customErr.code
	}

	return ECUnknown
}
