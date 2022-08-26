package httpclient

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/DWHengr/aurora/pkg/logger"
	"github.com/DWHengr/aurora/pkg/misc/header"
	"io/ioutil"
	"net/http"
)

// MarshalNotHtml not encode html
func MarshalNotHtml(data interface{}) ([]byte, error) {
	bf := bytes.NewBuffer([]byte{})
	jsonEncoder := json.NewEncoder(bf)
	jsonEncoder.SetEscapeHTML(false)
	if err := jsonEncoder.Encode(data); err != nil {
		return nil, err
	}
	return bf.Bytes(), nil
}

// POST http post
func POST(ctx context.Context, client *http.Client, uri string, params interface{}, entity interface{}) error {
	paramByte, err := MarshalNotHtml(params)
	if err != nil {
		return err
	}

	fmt.Println("Request-Id:" + header.GetRequestId(ctx) + " http request uri:" + uri + "http request params:" + string(paramByte))
	reader := bytes.NewReader(paramByte)
	req, err := http.NewRequest("POST", uri, reader)
	if err != nil {
		return err
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add(header.GetRequestIDKV(ctx).Wreck())

	response, err := client.Do(req)
	if err != nil {
		logger.Logger.Errorw(err.Error(), header.GetRequestId(ctx))
		return err
	}
	defer response.Body.Close()
	err = DecomposeResp(response, entity)
	if err != nil {
		logger.Logger.Errorw(err.Error(), header.GetRequestId(ctx))
		return err
	}
	return err
}

// Request http
func Request(uri string, method string, params interface{}, entity interface{}, headers map[string]string) error {
	client = GetHttpClient()
	paramByte, err := json.Marshal(params)
	if err != nil {
		return err
	}
	reader := bytes.NewReader(paramByte)
	req, err := http.NewRequest(method, uri, reader)
	if err != nil {
		logger.Logger.Errorw(err.Error())
		return err
	}

	if headers != nil {
		for k, v := range headers {
			req.Header.Add(k, v)
		}
	}

	response, err := client.Do(req)
	if err != nil {
		logger.Logger.Errorw(err.Error())
		return err
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		logger.Logger.Errorw(err.Error())
		return err
	}
	if entity != nil {
		err = json.Unmarshal(body, entity)
	}
	if err != nil {
		logger.Logger.Errorw(err.Error())
		return err
	}
	return nil
}
