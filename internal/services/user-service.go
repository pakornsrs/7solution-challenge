package services

import (
	"context"
	"pakornssn/7solution-challenge/internal/constants"
	"pakornssn/7solution-challenge/internal/models"
	"pakornssn/7solution-challenge/internal/repositories"
	errorutil "pakornssn/7solution-challenge/pkg/error-util"
	"strings"
)

type IUserService interface {
	Register(ctx context.Context, request models.CreateUserRequest) (models.UserResponse, error)
}

type userService struct {
	userRepository repositories.IUserRepository
}

func NewUserService(userRepositoryClient repositories.IUserRepository) IUserService {
	return &userService{
		userRepository: userRepositoryClient,
	}
}

func (service *userService) Register(ctx context.Context, request models.CreateUserRequest) (models.UserResponse, error) {

	// check existed
	filter := models.GetUserFilterRequest{
		Email: strings.TrimSpace(request.Email),
	}

	user, _, err := service.userRepository.FindUser(ctx, filter)
	if err != nil {
		return models.UserResponse{}, err
	}

	if len(user) > 0 {
		errorResp := errorutil.GetConflictError(&constants.UserAlreeadyExistError)
		return models.UserResponse{}, &errorResp
	}

	// create user
	userDb, err := request.ToUserDb()
	if err != nil {
		return models.UserResponse{}, err
	}

	err = service.userRepository.CreateUser(ctx, userDb)
	if err != nil {
		return models.UserResponse{}, err
	}

	return userDb.ToUserResponse(), nil
}
