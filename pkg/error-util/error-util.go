package errorutil

import (
	"pakornssn/7solution-challenge/internal/constants"
)

type ServiceError struct {
	StatusCode      int
	ErrorCode       string
	Message         string
	OriginalMessage string
}

func (e *ServiceError) Error() string {
	return e.Message
}

func GetInternalServerError(originalError error, customError *constants.CustomError) ServiceError {
	if customError == nil {
		customError = &constants.UnknownError
	}
	return ServiceError{
		StatusCode:      500,
		ErrorCode:       customError.Code,
		Message:         customError.Message,
		OriginalMessage: originalError.Error(),
	}
}
