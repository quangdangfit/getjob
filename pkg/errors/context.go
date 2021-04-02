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
			errorType:    customErr.errorType,
			wrappedError: customErr.wrappedError,
			context:      context,
		}
	}

	return CustomError{
		errorType:    Unknown,
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

// GetCode returns the error type
func GetType(err error) ErrorType {
	if customErr, ok := err.(CustomError); ok {
		return customErr.errorType
	}

	return Unknown
}
