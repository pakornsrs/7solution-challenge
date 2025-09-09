package validators

import (
	"errors"
	"pakornssn/7solution-challenge/internal/models"
	"regexp"
	"strings"
)

type ValidateResult struct {
	IsPass      bool
	ErrorDetail error
}

var emailRe = regexp.MustCompile(`^[A-Za-z0-9._%+-]+@[A-Za-z0-9.-]+\.[A-Za-z]{2,}$`)

func CreateUserRequestValidator(request models.CreateUserRequest) ValidateResult {
	if len(strings.TrimSpace(request.Name)) == 0 {
		return ValidateResult{IsPass: false, ErrorDetail: errors.New("name must not be empty")}
	}

	if len(strings.TrimSpace(request.Email)) == 0 {
		return ValidateResult{IsPass: false, ErrorDetail: errors.New("email must not be empty")}
	}

	if !emailRe.MatchString(request.Email) {
		return ValidateResult{IsPass: false, ErrorDetail: errors.New("email format is incorrect")}
	}

	if len(strings.TrimSpace(request.Password)) < 8 {
		return ValidateResult{IsPass: false, ErrorDetail: errors.New("password must not be empty and length must greater than 8 digit")}
	}

	return ValidateResult{IsPass: true}
}
