package api

import (
	"github.com/DWHengr/aurora/internal/alertcore"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Prometheus struct {
}

func prometheusRouter(engine *gin.Engine) {
	prometheus := NewPrometheus()
	engine.POST("/api/v2/alerts", prometheus.Alerts)
}

func NewPrometheus() *Prometheus {
	return &Prometheus{}
}

type PrometheusReq struct {
	Labels      map[string]string `json:"labels"`
	Annotations map[string]string `json:"annotations"`
}

func (p *Prometheus) Alerts(c *gin.Context) {
	reqs := &[]*PrometheusReq{}
	if err := c.ShouldBind(reqs); err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	for _, req := range *reqs {
		msg := &alertcore.AlertMessage{
			Name:      req.Labels["alertname"],
			Value:     req.Annotations["value"],
			Summary:   req.Annotations["summary"],
			UniqueId:  req.Labels["uniqueid"],
			Attribute: make(map[string]interface{}),
		}
		for k, v := range req.Labels {
			msg.Attribute[k] = v
		}
		for k, v := range req.Annotations {
			msg.Attribute[k] = v
		}
		alertcore.GetAlerterSingle().Receive(msg)
	}
}
