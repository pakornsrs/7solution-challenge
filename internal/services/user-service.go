package services

import (
	"context"
	"pakornssn/7solution-challenge/internal/constants"
	"pakornssn/7solution-challenge/internal/models"
	"pakornssn/7solution-challenge/internal/repositories"
	errorutil "pakornssn/7solution-challenge/pkg/error-util"
	"strings"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type IUserService interface {
	Register(ctx context.Context, request models.CreateUserRequest) (models.UserResponse, error)
	GetUserById(ctx context.Context, userId string) (models.UserResponse, error)
	GetAllUser(ctx context.Context, pagination *models.PaginationRequest) (models.GetAllUserResponse, error)
	UpdateUser(ctx context.Context, request models.UpdateUserRequest) (models.UpdateUserResponse, error)
	DeleteUser(ctx context.Context, userId string) error
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

	userDb.Id = primitive.NewObjectID()

	err = service.userRepository.CreateUser(ctx, userDb)
	if err != nil {
		return models.UserResponse{}, err
	}

	return userDb.ToUserResponse(), nil
}

func (service *userService) GetUserById(ctx context.Context, userId string) (models.UserResponse, error) {

	filter := models.GetUserFilterRequest{
		UserId: userId,
	}

	users, _, err := service.userRepository.FindUser(ctx, filter)
	if err != nil {
		return models.UserResponse{}, err
	}

	if len(users) == 0 {
		errorResp := errorutil.GetServerErrorResponse(404, nil, &constants.UserNotFoundError)
		return models.UserResponse{}, &errorResp
	}

	return users[0].ToUserResponse(), nil
}

func (service *userService) GetAllUser(ctx context.Context, pagination *models.PaginationRequest) (models.GetAllUserResponse, error) {

	filter := models.GetUserFilterRequest{}

	if pagination != nil {
		filter.Pagination = pagination
	}

	users, paginationResp, err := service.userRepository.FindUser(ctx, filter)
	if err != nil {
		return models.GetAllUserResponse{}, err
	}

	listUser := []models.UserResponse{}
	for _, user := range users {
		listUser = append(listUser, user.ToUserResponse())
	}

	resp := models.GetAllUserResponse{
		List:       listUser,
		Pagination: paginationResp,
	}

	return resp, nil
}

func (service *userService) UpdateUser(ctx context.Context, request models.UpdateUserRequest) (models.UpdateUserResponse, error) {

	// check existed
	if len(request.UpdatedEmail) > 0 {
		filter := models.GetUserFilterRequest{
			Email: strings.TrimSpace(request.UpdatedEmail),
		}

		user, _, err := service.userRepository.FindUser(ctx, filter)
		if err != nil {
			return models.UpdateUserResponse{}, err
		}

		if len(user) > 0 {
			errorResp := errorutil.GetConflictError(&constants.UserAlreeadyExistError)
			return models.UpdateUserResponse{}, &errorResp
		}
	}

	// update user
	updatedRequest, err := request.ToUpdateUserDb()
	if err != nil {
		return models.UpdateUserResponse{}, err
	}
	err = service.userRepository.UpdateUser(ctx, updatedRequest)
	if err != nil {
		return models.UpdateUserResponse{}, err
	}

	// get updated user
	user, err := service.GetUserById(ctx, request.UserId)
	if err != nil {
		return models.UpdateUserResponse{}, err
	}

	return models.UpdateUserResponse{UpdatedUser: user}, nil
}

func (service *userService) DeleteUser(ctx context.Context, userId string) error {
	// check existed
	filter := models.GetUserFilterRequest{
		UserId: strings.TrimSpace(userId),
	}

	user, _, err := service.userRepository.FindUser(ctx, filter)
	if err != nil {
		return err
	}

	if len(user) == 0 {
		errorResp := errorutil.GetServerErrorResponse(404, nil, &constants.UserNotFoundError)
		return &errorResp
	}

	// delete user
	primitiveId, err := primitive.ObjectIDFromHex(userId)
	if err != nil {
		respError := errorutil.GetServerErrorResponse(500, err, &constants.UserIdFormatIncorrectError)
		return &respError
	}

	err = service.userRepository.DeleteUser(ctx, primitiveId)
	if err != nil {
		return err
	}

	return nil
}
