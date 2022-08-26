package httpclient

import (
	"encoding/json"
	"errors"
	"github.com/DWHengr/aurora/pkg/logger"
	"github.com/DWHengr/aurora/pkg/misc/header"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
)

const (
	Unknown = -1
	Success = 0
	Error   = 1
)

type R struct {
	err  error
	Code int64       `json:"code"`
	Msg  string      `json:"msg,omitempty"`
	Data interface{} `json:"data"`
}

func DecomposeResp(response *http.Response, entity interface{}) error {
	if response.StatusCode != http.StatusOK {
		return errors.New("request error")
	}

	r := &R{}
	r.Data = entity

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(body, r)
	if err != nil {
		return err
	}
	return nil
}

type resp interface{}

func Format(resp resp, err error) *R {
	if err == nil {
		return &R{
			Code: Success,
			Data: resp,
		}
	}
	return &R{
		Code: Error,
		Msg:  err.Error(),
	}
}

func (r *R) Context(c *gin.Context, code ...int) {
	status := http.StatusOK
	if r.Code == Unknown {
		status = http.StatusInternalServerError
		if r.err != nil {
			logger.Logger.Errorw(r.err.Error(), header.GINRequestID(c))
		} else {
			logger.Logger.Errorw("unknown err", header.GINRequestID(c))
		}
	} else if len(code) != 0 {
		status = code[0]
	}
	c.JSON(status, r)
}
