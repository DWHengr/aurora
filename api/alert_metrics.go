package api

import (
	"github.com/DWHengr/aurora/internal/models"
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
		httpclient.Format(nil, err).Context(c)
	}
	httpclient.Format(resp, err).Context(c)
}
