package api

import (
	"github.com/DWHengr/aurora/internal/models"
	"github.com/DWHengr/aurora/internal/service"
	"github.com/DWHengr/aurora/pkg/httpclient"
	"github.com/DWHengr/aurora/pkg/logger"
	"github.com/gin-gonic/gin"
	"net/http"
)

type AlertSilences struct {
	alertSilencesService service.AlertSilencesService
}

func alertSilenceRouter(engine *gin.Engine) {
	alertRules := NewAlertSilences()
	group := engine.Group("/api/v1/silence")
	group.POST("/create", alertRules.CreateSilence)
}

func NewAlertSilences() *AlertSilences {
	alertSilencesService, _ := service.NewAlertSilencesService()
	return &AlertSilences{
		alertSilencesService: alertSilencesService,
	}
}

func (a *AlertSilences) CreateSilence(c *gin.Context) {
	reqs := &models.AlertSilences{}
	if err := c.ShouldBind(reqs); err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		httpclient.Format(nil, err).Context(c)
		return
	}
	resp, err := a.alertSilencesService.Create(reqs)
	if err != nil {
		logger.Logger.Error(err)
		httpclient.Format(nil, err).Context(c)
	}
	httpclient.Format(resp, err).Context(c)
}
