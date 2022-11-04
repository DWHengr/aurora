package api

import (
	"errors"
	"github.com/DWHengr/aurora/pkg/config"
	"github.com/DWHengr/aurora/pkg/httpclient"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Login(c *gin.Context) {
	user := &Users{}
	if err := c.ShouldBind(user); err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		httpclient.Format(nil, err).Context(c)
		return
	}
	allConfig, _ := config.GetAllConfig()
	if user.Username != allConfig.Aurora.Username || user.Password != allConfig.Aurora.Password {
		httpclient.Format(nil, errors.New("ERROR Incorrect username or password")).Context(c)
		return
	}
	token, err := GenToken(*user)
	httpclient.Format(token, err).Context(c)
}
