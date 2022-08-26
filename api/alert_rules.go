package api

import (
	"github.com/DWHengr/aurora/internal/models"
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
