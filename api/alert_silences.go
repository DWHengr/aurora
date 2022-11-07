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

type AlertSilences struct {
	alertSilencesService service.AlertSilencesService
}

func alertSilenceRouter(engine *gin.Engine) {
	alertSilence := NewAlertSilences()
	group := engine.Group("/api/v1/silence")
	group.POST("/create", alertSilence.CreateSilence)
	group.POST("/deletes", alertSilence.DeletesSilence)
	group.POST("/update", alertSilence.UpdateSilence)
	group.POST("/page", alertSilence.PageSilence)
	group.POST("/all", alertSilence.GetAllSilences)
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
	}
	httpclient.Format(resp, err).Context(c)
}

func (a *AlertSilences) DeletesSilence(c *gin.Context) {
	ids := &[]string{}
	if err := c.ShouldBind(ids); err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		httpclient.Format(nil, err).Context(c)
		return
	}
	err := a.alertSilencesService.Deletes(*ids)
	if err != nil {
		logger.Logger.Error(err)
		httpclient.Format(nil, err).Context(c)
		return
	}
	httpclient.Format("delete success", nil).Context(c)
}

func (a *AlertSilences) UpdateSilence(c *gin.Context) {
	reqs := &models.AlertSilences{}
	if err := c.ShouldBind(reqs); err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		httpclient.Format(nil, err).Context(c)
		return
	}
	resp, err := a.alertSilencesService.Update(reqs)
	if err != nil {
		logger.Logger.Error(err)
	}
	httpclient.Format(resp, err).Context(c)
}

func (a *AlertSilences) GetAllSilences(c *gin.Context) {
	resp, err := a.alertSilencesService.GetAllAlertSilences()
	if err != nil {
		logger.Logger.Error(err)
	}
	httpclient.Format(resp, err).Context(c)
}

func (a *AlertSilences) PageSilence(c *gin.Context) {
	page := &page.ReqPage{}
	if err := c.ShouldBind(page); err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		httpclient.Format(nil, err).Context(c)
		return
	}
	resp, err := a.alertSilencesService.Page(page)
	if err != nil {
		logger.Logger.Error(err)
	}
	httpclient.Format(resp, err).Context(c)
}
