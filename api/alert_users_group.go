package api

import (
	"github.com/DWHengr/aurora/internal/models"
	"github.com/DWHengr/aurora/internal/page"
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
	group.POST("/page", alertUsersGroup.PageUserGroup)
	group.POST("/update", alertUsersGroup.UpdateUserGroup)
	group.POST("/deletes", alertUsersGroup.DeletesUserGroup)
	group.POST("/all", alertUsersGroup.AllUserGroup)
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
	}
	httpclient.Format(resp, err).Context(c)
}

func (a *AlertUsersGroup) DeletesUserGroup(c *gin.Context) {
	ids := &[]string{}
	if err := c.ShouldBind(ids); err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		httpclient.Format(nil, err).Context(c)
		return
	}
	err := a.alertUsersGroupService.Deletes(*ids)
	if err != nil {
		logger.Logger.Error(err)
		httpclient.Format(nil, err).Context(c)
		return
	}
	httpclient.Format("delete success", nil).Context(c)
}

func (a *AlertUsersGroup) UpdateUserGroup(c *gin.Context) {
	reqs := &models.AlertUsersGroup{}
	if err := c.ShouldBind(reqs); err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		httpclient.Format(nil, err).Context(c)
		return
	}
	resp, err := a.alertUsersGroupService.Update(reqs)
	if err != nil {
		logger.Logger.Error(err)
	}
	httpclient.Format(resp, err).Context(c)
}

func (a *AlertUsersGroup) PageUserGroup(c *gin.Context) {
	page := &page.ReqPage{}
	if err := c.ShouldBind(page); err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		httpclient.Format(nil, err).Context(c)
		return
	}
	resp, err := a.alertUsersGroupService.Page(page)
	if err != nil {
		logger.Logger.Error(err)
	}
	httpclient.Format(resp, err).Context(c)
}

func (a *AlertUsersGroup) AllUserGroup(c *gin.Context) {
	resp, err := a.alertUsersGroupService.All()
	if err != nil {
		logger.Logger.Error(err)
	}
	httpclient.Format(resp, err).Context(c)
}
