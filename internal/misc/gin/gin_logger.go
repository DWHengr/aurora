package gin

import (
	"net"
	"net/http"
	"net/http/httputil"
	"os"
	"runtime/debug"
	"strings"
	"time"

	pkglogger "aurora/internal/logger"
	"aurora/internal/misc/header"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// GetRequestID get Request-Id from header
func GetRequestID(c *gin.Context) zap.Field {
	return ginHeader(c, header.RequestID)
}

// GetTimezone get Timezone from header
func GetTimezone(c *gin.Context) zap.Field {
	return ginHeader(c, header.Timezone)
}

func ginHeader(c *gin.Context, key string) zap.Field {
	val := c.GetHeader(key)
	return zap.String(key, val)
}

// LoggerFunc gin logger
func LoggerFunc() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Start timer
		start := time.Now()
		path := c.Request.URL.Path
		raw := c.Request.URL.RawQuery

		// Process request
		c.Next()

		if raw != "" {
			path = path + "?" + raw
		}

		// Stop timer
		pkglogger.Logger.Infow("[GIN]",
			GetRequestID(c),
			zap.Float64("latency", time.Now().Sub(start).Seconds()),
			zap.String("clientIP", c.ClientIP()),
			zap.String("method", c.Request.Method),
			zap.Int("statusCode", c.Writer.Status()),
			zap.String("errorMessage", c.Errors.ByType(gin.ErrorTypePrivate).String()),
			zap.Int("bodySize", c.Writer.Size()),
			zap.String("path", path),
		)
	}
}

// RecoveryFunc gin recovery
func RecoveryFunc() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				var brokenPipe bool
				if ne, ok := err.(*net.OpError); ok {
					if se, ok := ne.Err.(*os.SyscallError); ok {
						if strings.Contains(strings.ToLower(se.Error()), "broken pipe") || strings.Contains(strings.ToLower(se.Error()), "connection reset by peer") {
							brokenPipe = true
						}
					}
				}

				httpRequest, _ := httputil.DumpRequest(c.Request, false)
				if brokenPipe {
					pkglogger.Logger.Error(c.Request.URL.Path,
						GetRequestID(c),
						zap.Any("error", err),
						zap.String("request", string(httpRequest)),
					)
					c.Error(err.(error))
					c.Abort()
					return
				}

				pkglogger.Logger.Error("[Recovery from panic]",
					GetRequestID(c),
					zap.Any("error", err),
					zap.String("request", string(httpRequest)),
					zap.String("stack", string(debug.Stack())),
				)

				c.AbortWithStatus(http.StatusInternalServerError)
			}
		}()
		c.Next()
	}
}
