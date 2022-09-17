package api

import (
	"github.com/DWHengr/aurora/internal/models"
	"github.com/DWHengr/aurora/internal/service"
	"github.com/DWHengr/aurora/pkg/httpclient"
	"github.com/DWHengr/aurora/pkg/logger"
	"github.com/gin-gonic/gin"
	"net/http"
)

type AlertUsers struct {
	alertUsersService service.AlertUsersService
}

func alertUsersRouter(engine *gin.Engine) {
	alertUsers := NewUserRules()
	group := engine.Group("/api/v1/user")
	group.POST("/create", alertUsers.CreateUser)
	group.POST("/deletes", alertUsers.DeletesUser)
	group.POST("/update", alertUsers.UpdateUser)
}

func NewUserRules() *AlertUsers {
	alertUsersService, _ := service.NewAlertUsersService()
	return &AlertUsers{
		alertUsersService: alertUsersService,
	}
}

func (a *AlertUsers) CreateUser(c *gin.Context) {
	reqs := &models.AlertUsers{}
	if err := c.ShouldBind(reqs); err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		httpclient.Format(nil, err).Context(c)
		return
	}
	resp, err := a.alertUsersService.Create(reqs)
	if err != nil {
		logger.Logger.Error(err)
		httpclient.Format(nil, err).Context(c)
	}
	httpclient.Format(resp, err).Context(c)
}

func (a *AlertUsers) DeletesUser(c *gin.Context) {
	ids := &[]string{}
	if err := c.ShouldBind(ids); err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		httpclient.Format(nil, err).Context(c)
		return
	}
	err := a.alertUsersService.Deletes(*ids)
	if err != nil {
		logger.Logger.Error(err)
		httpclient.Format(nil, err).Context(c)
		return
	}
	httpclient.Format("delete success", nil).Context(c)
}

func (a *AlertUsers) UpdateUser(c *gin.Context) {
	reqs := &models.AlertUsers{}
	if err := c.ShouldBind(reqs); err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		httpclient.Format(nil, err).Context(c)
		return
	}
	resp, err := a.alertUsersService.Update(reqs)
	if err != nil {
		logger.Logger.Error(err)
		httpclient.Format(nil, err).Context(c)
	}
	httpclient.Format(resp, err).Context(c)
}
