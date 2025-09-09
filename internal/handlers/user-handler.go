package handlers

import (
	"context"
	"pakornssn/7solution-challenge/internal/constants"
	"pakornssn/7solution-challenge/internal/models"
	"pakornssn/7solution-challenge/internal/services"
	"pakornssn/7solution-challenge/internal/validators"
	errorutil "pakornssn/7solution-challenge/pkg/error-util"
	httputil "pakornssn/7solution-challenge/pkg/http-util"
	"time"

	"github.com/gin-gonic/gin"
)

type IUserHandler interface {
	Register(gin *gin.Context)
}

type userHandler struct {
	userService services.IUserService
}

func NewUserHandler(userServiceClient services.IUserService) IUserHandler {
	return &userHandler{
		userService: userServiceClient,
	}
}

func (handler *userHandler) Register(gin *gin.Context) {
	request := models.CreateUserRequest{}

	if err := gin.ShouldBindJSON(&request); err != nil {
		httputil.HttpBadRequestResponse(gin, constants.BadRequestError.Message, err.Error())
		return
	}

	validateResult := validators.CreateUserRequestValidator(request)

	if !validateResult.IsPass {
		errResp := errorutil.GetServerErrorResponse(422, validateResult.ErrorDetail, &constants.ValidateRequestError)
		httputil.HttpErrorResponse(gin, &errResp)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	resp, err := handler.userService.Register(ctx, request)
	if err != nil {
		httputil.HttpErrorResponse(gin, err)
		return
	}

	httputil.HttpSuccessResponse(gin, resp)
}
