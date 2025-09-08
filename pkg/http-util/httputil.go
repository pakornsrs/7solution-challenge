package httputil

import (
	"net/http"
	"pakornssn/7solution-challenge/internal/constants"
	"pakornssn/7solution-challenge/internal/models"

	"github.com/gin-gonic/gin"
)

func ResponseSuccessStatusWithBody[T any](gin *gin.Context, resp T) {
	gin.JSON(
		http.StatusOK,
		models.CreateSuccessResponseModel(
			resp,
		),
	)
}

func ResponseBadRequest[T any](gin *gin.Context, resp T) {
	gin.JSON(
		http.StatusBadRequest,
		models.CreateResponseModel(
			resp,
			constants.BadRequestErrorCode,
			constants.BadRequestErrorMessage,
		),
	)
}

func ResponseServiceInternalError[T any](gin *gin.Context, resp T) {
	gin.JSON(
		http.StatusInternalServerError,
		models.CreateResponseModel(
			resp,
			constants.InternalProcessFailedErrorCode,
			constants.InternalProcessFailedErrorMessage,
		),
	)
}
