package api

import (
	"errors"
	"github.com/DWHengr/aurora/internal/models"
	"github.com/DWHengr/aurora/internal/page"
	"github.com/DWHengr/aurora/internal/service"
	"github.com/DWHengr/aurora/pkg/httpclient"
	"github.com/DWHengr/aurora/pkg/logger"
	"github.com/gin-gonic/gin"
	"net/http"
)

type AlertRules struct {
	alertRulesService service.AlertRulesService
}

func alertRuleRouter(engine *gin.Engine) {
	alertRules := NewAlertRules()
	group := engine.Group("/api/v1/rule")
	group.POST("/create", alertRules.CreateRule)
	group.POST("/update", alertRules.UpdateRule)
	group.POST("/update/status", alertRules.UpdateRuleStatus)
	group.POST("/page", alertRules.PageRule)
	group.POST("/delete/:id", alertRules.DeleteRule)
	group.POST("/deletes", alertRules.DeletesRule)
	group.POST("/details/:id", alertRules.DetailsRule)
}

func NewAlertRules() *AlertRules {
	alertRulesService, _ := service.NewAlertRulesService()
	return &AlertRules{
		alertRulesService: alertRulesService,
	}
}

func (a *AlertRules) CreateRule(c *gin.Context) {
	reqs := &models.AlertRules{}
	if err := c.ShouldBind(reqs); err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		httpclient.Format(nil, err).Context(c)
		return
	}
	resp, err := a.alertRulesService.Create(reqs)
	if err != nil {
		logger.Logger.Error(err)
	}
	httpclient.Format(resp, err).Context(c)
}

func (a *AlertRules) DeleteRule(c *gin.Context) {
	ruleId, ok := c.Params.Get("id")
	if !ok {
		httpclient.Format(nil, errors.New("invalid URI")).Context(c)
		return
	}
	err := a.alertRulesService.Delete(ruleId)
	if err != nil {
		logger.Logger.Error(err)
		httpclient.Format(nil, err).Context(c)
		return
	}
	httpclient.Format("delete success", nil).Context(c)
}

func (a *AlertRules) DetailsRule(c *gin.Context) {
	ruleId, ok := c.Params.Get("id")
	if !ok {
		httpclient.Format(nil, errors.New("invalid URI")).Context(c)
		return
	}
	resp, err := a.alertRulesService.Details(ruleId)
	if err != nil {
		logger.Logger.Error(err)
	}
	httpclient.Format(resp, err).Context(c)
}

func (a *AlertRules) UpdateRule(c *gin.Context) {
	reqs := &models.AlertRules{}
	if err := c.ShouldBind(reqs); err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		httpclient.Format(nil, err).Context(c)
		return
	}
	resp, err := a.alertRulesService.Update(reqs)
	if err != nil {
		logger.Logger.Error(err)
	}
	httpclient.Format(resp, err).Context(c)
}

func (a *AlertRules) DeletesRule(c *gin.Context) {
	ids := &[]string{}
	if err := c.ShouldBind(ids); err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		httpclient.Format(nil, err).Context(c)
		return
	}
	err := a.alertRulesService.Deletes(*ids)
	if err != nil {
		logger.Logger.Error(err)
		httpclient.Format(nil, err).Context(c)
		return
	}
	httpclient.Format("delete success", nil).Context(c)
}

func (a *AlertRules) PageRule(c *gin.Context) {
	page := &page.ReqPage{}
	if err := c.ShouldBind(page); err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		httpclient.Format(nil, err).Context(c)
		return
	}
	resp, err := a.alertRulesService.Page(page)
	if err != nil {
		logger.Logger.Error(err)
	}
	httpclient.Format(resp, err).Context(c)
}

func (a *AlertRules) UpdateRuleStatus(c *gin.Context) {
	reqs := &models.AlertRules{}
	if err := c.ShouldBind(reqs); err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		httpclient.Format(nil, err).Context(c)
		return
	}
	resp, err := a.alertRulesService.UpdateStatus(reqs)
	if err != nil {
		logger.Logger.Error(err)
	}
	httpclient.Format(resp, err).Context(c)
}
