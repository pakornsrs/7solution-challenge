package routes

import (
	"fmt"
	"os"
	httputil "pakornssn/7solution-challenge/pkg/http-util"

	"github.com/gin-gonic/gin"
)

func SetupServiceHealthCheckRoute(router *gin.Engine) {

	base := "/api/user"

	router.GET(fmt.Sprintf("%s/health-check", base), healthCheck)
}

func healthCheck(c *gin.Context) {
	resp := healthCheckResponse{
		Environment: os.Getenv("ENV"),
		Version:     os.Getenv("VERSION"),
	}
	httputil.ResponseSuccessStatusWithBody(c, resp)
}

type healthCheckResponse struct {
	Environment string `json:"environment"`
	Version     string `json:"version"`
}
