package httputil

import (
	"net/http"
	"pakornssn/7solution-challenge/internal/constants"
	"pakornssn/7solution-challenge/internal/models"
	errorutil "pakornssn/7solution-challenge/pkg/error-util"

	"github.com/gin-gonic/gin"
)

func HttpSuccessResponse(gin *gin.Context, resp interface{}) {
	gin.JSON(http.StatusOK, models.CreateResponseModel(resp, constants.SuccessCode, constants.SuccessMessage, constants.SuccessMessage))
}

func HttpErrorResponse(gin *gin.Context, errorDetail error) {
	var resp interface{}
	httpStatus := http.StatusInternalServerError
	errorCode := constants.InternalProcessFailedErrorCode
	errorMessage := ""
	originalErrorMessage := ""

	if appErr, ok := errorDetail.(*errorutil.ServiceError); ok {
		httpStatus = appErr.StatusCode
		errorCode = appErr.ErrorCode
		errorMessage = appErr.Message
		originalErrorMessage = appErr.OriginalMessage
	} else {
		errorMessage = errorDetail.Error()
	}

	gin.JSON(httpStatus, models.CreateResponseModel(resp, errorCode, errorMessage, originalErrorMessage))
}

func HttpBadRequestResponse(gin *gin.Context, message, orriginalMessage string) {
	gin.JSON(http.StatusBadRequest, models.CreateResponseModel[any](nil, constants.BadRequestErrorCode, message, orriginalMessage))
}
