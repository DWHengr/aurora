package api

import (
	"errors"
	"github.com/DWHengr/aurora/internal/page"
	"github.com/DWHengr/aurora/internal/service"
	"github.com/DWHengr/aurora/pkg/httpclient"
	"github.com/DWHengr/aurora/pkg/logger"
	"github.com/gin-gonic/gin"
	"net/http"
)

type AlertRecords struct {
	alertRecordsService service.AlertRecordsService
}

func alertRecordsRouter(engine *gin.Engine) {
	alertMetrics := NewAlertRecords()
	group := engine.Group("/api/v1/record")
	group.POST("/delete/:id", alertMetrics.DeleteRecord)
	group.POST("/page", alertMetrics.PageRecord)
}

func NewAlertRecords() *AlertRecords {
	alertRecordsService, _ := service.NewAlertRecordsService()
	return &AlertRecords{
		alertRecordsService: alertRecordsService,
	}
}

func (a *AlertRecords) DeleteRecord(c *gin.Context) {
	ruleId, ok := c.Params.Get("id")
	if !ok {
		httpclient.Format(nil, errors.New("invalid URI")).Context(c)
		return
	}
	err := a.alertRecordsService.Delete(ruleId)
	if err != nil {
		logger.Logger.Error(err)
		httpclient.Format(nil, err).Context(c)
		return
	}
	httpclient.Format("delete success", nil).Context(c)
}

func (a *AlertRecords) PageRecord(c *gin.Context) {
	page := &page.ReqPage{}
	if err := c.ShouldBind(page); err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		httpclient.Format(nil, err).Context(c)
		return
	}
	resp, err := a.alertRecordsService.Page(page)
	if err != nil {
		logger.Logger.Error(err)
	}
	httpclient.Format(resp, err).Context(c)
}
