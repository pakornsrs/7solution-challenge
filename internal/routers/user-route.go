package routes

import (
	"pakornssn/7solution-challenge/internal/handlers"
	"pakornssn/7solution-challenge/internal/middleware"

	"github.com/gin-gonic/gin"
)

func SetupUserServiceRoute(router *gin.Engine, userHandler handlers.IUserHandler) {

	publicPath := router.Group("/api/user")

	securePath := publicPath.Group("/")
	securePath.Use(middleware.ValidateToken())

	publicPath.POST("/register", userHandler.Register)
	securePath.GET("/:userid", userHandler.GetUserById)
	securePath.GET("/all", userHandler.GetAllUser)
	securePath.PUT("/update", userHandler.UpdateUser)
}
