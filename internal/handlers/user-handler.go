package handlers

import (
	"context"
	"pakornssn/7solution-challenge/internal/constants"
	"pakornssn/7solution-challenge/internal/models"
	"pakornssn/7solution-challenge/internal/services"
	"pakornssn/7solution-challenge/internal/validators"
	errorutil "pakornssn/7solution-challenge/pkg/error-util"
	httputil "pakornssn/7solution-challenge/pkg/http-util"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

type IUserHandler interface {
	Register(gin *gin.Context)
	GetUserById(gin *gin.Context)
	GetAllUser(gin *gin.Context)
	UpdateUser(gin *gin.Context)
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

func (handler *userHandler) GetUserById(gin *gin.Context) {
	userId := strings.TrimSpace(gin.Param("userid"))

	validateResult := validators.ValidateUserId(userId)
	if !validateResult.IsPass {
		errResp := errorutil.GetServerErrorResponse(422, validateResult.ErrorDetail, &constants.ValidateRequestError)
		httputil.HttpErrorResponse(gin, &errResp)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	resp, err := handler.userService.GetUserById(ctx, userId)
	if err != nil {
		httputil.HttpErrorResponse(gin, err)
		return
	}

	httputil.HttpSuccessResponse(gin, resp)
}

func (handler *userHandler) GetAllUser(gin *gin.Context) {
	page := strings.TrimSpace(gin.Query("page"))
	itemPerPage := strings.TrimSpace(gin.Query("itemperpage"))

	var paginationRequest *models.PaginationRequest
	if len(page) > 0 && len(itemPerPage) > 0 {
		validateResult, pagination := validators.ValidatePagination(page, itemPerPage)
		if !validateResult.IsPass {
			errResp := errorutil.GetServerErrorResponse(422, validateResult.ErrorDetail, &constants.ValidateRequestError)
			httputil.HttpErrorResponse(gin, &errResp)
			return
		}
		paginationRequest = pagination
	} else {
		// no pagination
		paginationRequest = nil
	}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	resp, err := handler.userService.GetAllUser(ctx, paginationRequest)
	if err != nil {
		httputil.HttpErrorResponse(gin, err)
		return
	}

	httputil.HttpSuccessResponse(gin, resp)
}

func (handler *userHandler) UpdateUser(gin *gin.Context) {
	request := models.UpdateUserRequest{}

	if err := gin.ShouldBindJSON(&request); err != nil {
		httputil.HttpBadRequestResponse(gin, constants.BadRequestError.Message, err.Error())
		return
	}

	validateResult := validators.UpdateUserRequestValidator(request)
	if !validateResult.IsPass {
		errResp := errorutil.GetServerErrorResponse(422, validateResult.ErrorDetail, &constants.ValidateRequestError)
		httputil.HttpErrorResponse(gin, &errResp)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	resp, err := handler.userService.UpdateUser(ctx, request)
	if err != nil {
		httputil.HttpErrorResponse(gin, err)
		return
	}

	httputil.HttpSuccessResponse(gin, resp)
}
