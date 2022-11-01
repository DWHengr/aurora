package api

import (
	"github.com/DWHengr/aurora/pkg/httpclient"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Cors() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		method := ctx.Request.Method

		ctx.Header("Access-Control-Allow-Origin", "*")
		ctx.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token,Authorization,aurora-token")
		ctx.Header("Access-Control-Allow-Methods", "POST,GET,OPTIONS,DELETE,PATCH,PUT")
		ctx.Header("Access-Control-Expose-Headers", "Content-Length,Access-Control-Allow-Origin,Access-Control-Allow-Headers,Content-Type")
		ctx.Header("Access-Control-Allow-Credentials", "true")

		if method == "OPTIONS" {
			ctx.AbortWithStatus(http.StatusNoContent)
		}
	}
}
func JWTAuth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authHeader := ctx.Request.Header.Get("Aurora-Token")
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
