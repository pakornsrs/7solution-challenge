package services

import (
	"context"
	"pakornssn/7solution-challenge/internal/constants"
	"pakornssn/7solution-challenge/internal/models"
	"pakornssn/7solution-challenge/internal/repositories"
	authutil "pakornssn/7solution-challenge/pkg/auth-util"
	errorutil "pakornssn/7solution-challenge/pkg/error-util"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

type IAuthenticationService interface {
	AuthenticationUser(ctx context.Context, request models.AuthUserRequest) (models.AuthUserResponse, error)
}

type authenticationService struct {
	userRepository repositories.IUserRepository
}

func NewAuthenticationService(userRepositoryClient repositories.IUserRepository) IAuthenticationService {
	return &authenticationService{
		userRepository: userRepositoryClient,
	}
}

func (service *authenticationService) AuthenticationUser(ctx context.Context, request models.AuthUserRequest) (models.AuthUserResponse, error) {

	// find user
	filter := models.GetUserFilterRequest{
		Email: strings.TrimSpace(request.Email),
	}
	users, _, err := service.userRepository.FindUser(ctx, filter)
	if err != nil {
		return models.AuthUserResponse{}, err
	}

	if len(users) != 1 {
		errorResp := errorutil.GetServerErrorResponse(401, nil, &constants.AuthenticationError)
		return models.AuthUserResponse{}, &errorResp
	}

	user := users[0]

	// authentication password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password)); err != nil {
		errorResp := errorutil.GetServerErrorResponse(401, nil, &constants.AuthenticationError)
		return models.AuthUserResponse{}, &errorResp
	}

	token, err := authutil.GenerateAuthenticationToken(ctx, user.Id.Hex(), constants.TokenDuration)
	if err != nil {
		errorResp := errorutil.GetServerErrorResponse(500, err, &constants.UnknownError)
		return models.AuthUserResponse{}, &errorResp
	}

	userResponse := models.AuthUserResponse{
		User:  user.ToUserResponse(),
		Token: token,
	}

	return userResponse, nil

}
