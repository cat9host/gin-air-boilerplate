package response

import (
	"github.com/cat9host/gin-air-boilerplate/internal/interfaces"
)

func BadRequestError(message string) *interfaces.RestError {
	return &interfaces.RestError{
		Message: message,
		Error:   "Invalid Request",
		Code:    interfaces.InternalError,
	}
}
func CriticalError(message string) *interfaces.RestError {
	return &interfaces.RestError{
		Message: message,
		Error:   "Critical error",
		Code:    interfaces.CriticalError,
	}
}

func SpecificBadRequestError(message string, code interfaces.StatusCode) *interfaces.RestError {
	return &interfaces.RestError{
		Message: message,
		Error:   "Invalid Request",
		Code:    code,
	}
}

func UnauthorizedRequestError(message string) *interfaces.RestError {
	return &interfaces.RestError{
		Message: message,
		Error:   "Unauthorized Request",
		Code:    interfaces.AuthError,
	}
}
