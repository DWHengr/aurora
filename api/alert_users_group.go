package api

import (
	"github.com/DWHengr/aurora/internal/models"
	"github.com/DWHengr/aurora/internal/service"
	"github.com/DWHengr/aurora/pkg/httpclient"
	"github.com/DWHengr/aurora/pkg/logger"
	"github.com/gin-gonic/gin"
	"net/http"
)

type AlertUsersGroup struct {
	alertUsersGroupService service.AlertUsersGroupService
}

func alertUsersGroupRouter(engine *gin.Engine) {
	alertUsersGroup := NewUserRulesGroup()
	group := engine.Group("/api/v1/usergroup")
	group.POST("/create", alertUsersGroup.CreateUserGroup)
}

func NewUserRulesGroup() *AlertUsersGroup {
	alertUsersGroupService, _ := service.NewAlertUsersGroupService()
	return &AlertUsersGroup{
		alertUsersGroupService: alertUsersGroupService,
	}
}

func (a *AlertUsersGroup) CreateUserGroup(c *gin.Context) {
	reqs := &models.AlertUsersGroup{}
	if err := c.ShouldBind(reqs); err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		httpclient.Format(nil, err).Context(c)
		return
	}
	resp, err := a.alertUsersGroupService.Create(reqs)
	if err != nil {
		logger.Logger.Error(err)
		httpclient.Format(nil, err).Context(c)
	}
	httpclient.Format(resp, err).Context(c)
}
