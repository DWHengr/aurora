package httpclient

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
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
