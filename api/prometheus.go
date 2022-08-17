package api

import (
	"github.com/DWHengr/aurora/internal/alert"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Prometheus struct {
}

func NewPrometheus() *Prometheus {
	return &Prometheus{}
}

func (p *Prometheus) Alerts(c *gin.Context) {
	msgs := &[]*alert.AlertMessage{}
	if err := c.ShouldBind(msgs); err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	for _, msg := range *msgs {
		alert.AlertInstance.Receive(msg)
	}
}
