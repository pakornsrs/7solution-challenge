package routes

import (
	"fmt"
	"pakornssn/7solution-challenge/internal/handlers"

	"github.com/gin-gonic/gin"
)

func SetupUserServiceRoute(router *gin.Engine, userHandler handlers.IUserHandler) {

	base := "/api/user"

	router.POST(fmt.Sprintf("%s/register", base), userHandler.Register)
}
