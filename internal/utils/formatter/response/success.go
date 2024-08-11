package response

import (
	"github.com/cat9host/gin-air-boilerplate/internal/interfaces"
)

func GenericSuccess() *interfaces.RestSuccess {
	return &interfaces.RestSuccess{
		Code:   interfaces.Success,
		Result: nil,
	}
}

func SuccessWithResult[T any](result T) *interfaces.RestSuccess {
	return &interfaces.RestSuccess{
		Code:   interfaces.Success,
		Result: result,
	}
}
