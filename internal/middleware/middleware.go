package middleware

import (
	"pakornssn/7solution-challenge/internal/constants"
	authutil "pakornssn/7solution-challenge/pkg/auth-util"
	errorutil "pakornssn/7solution-challenge/pkg/error-util"
	httputil "pakornssn/7solution-challenge/pkg/http-util"

	"github.com/gin-gonic/gin"
)

func ValidateToken() gin.HandlerFunc {
	return func(gin *gin.Context) {
		token := gin.GetHeader("Authorization")
		if len(token) == 0 {
			errorResp := errorutil.GetServerErrorResponse(401, nil, &constants.UnauthorizedRequestError)
			httputil.HttpErrorResponse(gin, &errorResp)
			gin.Abort()
			return
		}

		claims, err := authutil.ParseAndValidateJWT(token)
		if err != nil {
			errorResp := errorutil.GetServerErrorResponse(401, err, &constants.UnauthorizedRequestError)
			httputil.HttpErrorResponse(gin, &errorResp)
			gin.Abort()
			return
		}

		gin.Set("userId", claims.Subject)
		gin.Set("audience", claims.Audience)
		gin.Set("issuer", claims.Issuer)

		gin.Next()
	}
}
