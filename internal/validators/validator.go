package validators

import (
	"errors"
	"pakornssn/7solution-challenge/internal/models"
	"regexp"
	"strconv"
	"strings"

	"go.mongodb.org/mongo-driver/bson/primitive"
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

func AuthenticationUserRequestValidator(request models.AuthUserRequest) ValidateResult {

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

func ValidateUserId(request string) ValidateResult {
	if len(strings.TrimSpace(request)) == 0 {
		return ValidateResult{IsPass: false, ErrorDetail: errors.New("userId must not be empty")}
	}

	if _, err := primitive.ObjectIDFromHex(request); err != nil {
		return ValidateResult{IsPass: false, ErrorDetail: errors.New("userId format incorrect")}
	}

	return ValidateResult{IsPass: true}
}

func ValidatePagination(page string, item string) (ValidateResult, *models.PaginationRequest) {
	currentPage, err := strconv.Atoi(page)
	if err != nil {
		return ValidateResult{IsPass: false, ErrorDetail: errors.New("current page format is incorrect")}, nil
	}

	itemPerPage, err := strconv.Atoi(item)
	if err != nil {
		return ValidateResult{IsPass: false, ErrorDetail: errors.New("item per page format is incorrect")}, nil
	}

	pagination := &models.PaginationRequest{
		ItemPerPage: int64(itemPerPage),
		CurrentPage: int64(currentPage),
	}

	return ValidateResult{IsPass: true}, pagination
}

func UpdateUserRequestValidator(request models.UpdateUserRequest) ValidateResult {
	if len(request.UpdatedName) != 0 {
		// prevent white space
		if len(strings.TrimSpace(request.UpdatedName)) == 0 {
			return ValidateResult{IsPass: false, ErrorDetail: errors.New("updated name must not be with space")}
		}
	}

	if len(strings.TrimSpace(request.UpdatedEmail)) != 0 {
		if !emailRe.MatchString(request.UpdatedEmail) {
			return ValidateResult{IsPass: false, ErrorDetail: errors.New("updated email format is incorrect")}
		}
	}

	if len(strings.TrimSpace(request.UserId)) == 0 {
		return ValidateResult{IsPass: false, ErrorDetail: errors.New("userId must not be empty")}
	}

	if _, err := primitive.ObjectIDFromHex(request.UserId); err != nil {
		return ValidateResult{IsPass: false, ErrorDetail: errors.New("userId format incorrect")}
	}

	return ValidateResult{IsPass: true}
}
