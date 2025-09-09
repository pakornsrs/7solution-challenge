package routes

import (
	"fmt"
	"pakornssn/7solution-challenge/internal/handlers"

	"github.com/gin-gonic/gin"
)

func SetupAuthenticationServiceRoute(router *gin.Engine, authHandler handlers.IAuthenticationHandler) {

	base := "/api/authentication"

	router.POST(fmt.Sprintf("%s/login", base), authHandler.Login)
}
