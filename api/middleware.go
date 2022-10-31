package api

import (
	"github.com/DWHengr/aurora/pkg/httpclient"
	"github.com/gin-gonic/gin"
	"net/http"
)

func JWTAuth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authHeader := ctx.Request.Header.Get("Access-Token")
		if authHeader == "" {
			ctx.JSON(http.StatusOK, gin.H{
				"code": httpclient.NoAccess,
				"msg":  "No access",
			})
			ctx.Abort()
			return
		}

		claims, err := ParseToken(authHeader)
		if err != nil {
			ctx.JSON(http.StatusOK, gin.H{
				"code": httpclient.TokenFailure,
				"msg":  "Invalid Token",
			})
			ctx.Abort()
			return
		}

		ctx.Set("claims", claims)
		ctx.Next()

	}
}
