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

type AlertMetrics struct {
	alertMetricsService service.AlertMetricsService
}

func alertMetricsRouter(engine *gin.Engine) {
	alertMetrics := NewMetricRules()
	group := engine.Group("/api/v1/metric")
	group.POST("/create", alertMetrics.CreateMetric)
	group.POST("/delete/:id", alertMetrics.DeleteMetric)
	group.POST("/page", alertMetrics.PageMetric)
	group.POST("/deletes", alertMetrics.DeletesMetric)
	group.POST("/update", alertMetrics.UpdateMetric)
}

func NewMetricRules() *AlertMetrics {
	alertMetricsService, _ := service.NewAlertMetricsService()
	return &AlertMetrics{
		alertMetricsService: alertMetricsService,
	}
}

func (a *AlertMetrics) CreateMetric(c *gin.Context) {
	reqs := &models.AlertMetrics{}
	if err := c.ShouldBind(reqs); err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		httpclient.Format(nil, err).Context(c)
		return
	}
	resp, err := a.alertMetricsService.Create(reqs)
	if err != nil {
		logger.Logger.Error(err)
	}
	httpclient.Format(resp, err).Context(c)
}

func (a *AlertMetrics) DeleteMetric(c *gin.Context) {
	ruleId, ok := c.Params.Get("id")
	if !ok {
		httpclient.Format(nil, errors.New("invalid URI")).Context(c)
		return
	}
	err := a.alertMetricsService.Delete(ruleId)
	if err != nil {
		logger.Logger.Error(err)
		httpclient.Format(nil, err).Context(c)
		return
	}
	httpclient.Format("delete success", nil).Context(c)
}

func (a *AlertMetrics) DeletesMetric(c *gin.Context) {
	ids := &[]string{}
	if err := c.ShouldBind(ids); err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		httpclient.Format(nil, err).Context(c)
		return
	}
	err := a.alertMetricsService.Deletes(*ids)
	if err != nil {
		logger.Logger.Error(err)
		httpclient.Format(nil, err).Context(c)
		return
	}
	httpclient.Format("delete success", nil).Context(c)
}

func (a *AlertMetrics) PageMetric(c *gin.Context) {
	page := &page.ReqPage{}
	if err := c.ShouldBind(page); err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		httpclient.Format(nil, err).Context(c)
		return
	}
	resp, err := a.alertMetricsService.Page(page)
	if err != nil {
		logger.Logger.Error(err)
	}
	httpclient.Format(resp, err).Context(c)
}

func (a *AlertMetrics) UpdateMetric(c *gin.Context) {
	reqs := &models.AlertMetrics{}
	if err := c.ShouldBind(reqs); err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		httpclient.Format(nil, err).Context(c)
		return
	}
	resp, err := a.alertMetricsService.Update(reqs)
	if err != nil {
		logger.Logger.Error(err)
	}
	httpclient.Format(resp, err).Context(c)
}
