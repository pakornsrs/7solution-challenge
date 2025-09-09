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

func GetServerErrorResponse(statusCode int, originalError error, customError *constants.CustomError) ServiceError {
	if customError == nil {
		customError = &constants.UnknownError
	}

	originalErrorMessage := ""
	if originalError != nil {
		originalErrorMessage = originalError.Error()
	}

	return ServiceError{
		StatusCode:      statusCode,
		ErrorCode:       customError.Code,
		Message:         customError.Message,
		OriginalMessage: originalErrorMessage,
	}
}

func GetConflictError(customError *constants.CustomError) ServiceError {
	if customError == nil {
		customError = &constants.UnknownError
	}
	return ServiceError{
		StatusCode:      409,
		ErrorCode:       customError.Code,
		Message:         customError.Message,
		OriginalMessage: "",
	}
}
