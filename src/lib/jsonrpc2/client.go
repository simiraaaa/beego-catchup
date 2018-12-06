package jsonrpc2

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/simiraaaa/beego-catchup/src/lib/httpclient"
	"github.com/simiraaaa/beego-catchup/src/lib/log"
)

// Client ... JSONRPC2のリクエストを行う
type Client struct {
	URL      string
	Headers  map[string]string
	Requests []*ClientRequest
}

// AddRequest ... JSONRPC2のリクエストを登録する
func (c *Client) AddRequest(id string, method string, params *json.RawMessage) {
	c.Requests = append(c.Requests, &ClientRequest{
		Version: version,
		ID:      id,
		Method:  method,
		Params:  params,
	})
}

// DoSingle ... JSONRPC2のシングルリクエストを行う
func (c *Client) DoSingle(ctx context.Context, method string, params interface{}) (*json.RawMessage, *ErrorResponse, error) {
	p, err := json.Marshal(params)
	if err != nil {
		log.Errorf(ctx, "json.Marchal error: %s", err.Error())
		return nil, nil, err
	}

	status, body, err := httpclient.PostJSON(ctx, c.URL, p, &httpclient.HTTPOption{Headers: c.Headers})
	if err != nil {
		log.Errorf(ctx, "httpclient.PostJSON error: %s", err.Error())
		return nil, nil, err
	}
	if status != http.StatusOK {
		err = fmt.Errorf("httpclient.PostJSON status: %d", status)
		log.Errorf(ctx, "%s", err.Error())
		return nil, nil, err
	}

	var res ClientResponse
	err = json.Unmarshal(body, &res)
	if err != nil {
		log.Errorf(ctx, "json.Unmarshal error: %s", err.Error())
		return nil, nil, err
	}

	return res.Result, res.Error, nil
}

// DoBatch ... JSONRPC2のバッチリクエストを行う
func (c *Client) DoBatch(ctx context.Context) ([]*ClientResponse, error) {
	p, err := json.Marshal(c.Requests)
	if err != nil {
		log.Errorf(ctx, "json.Marchal error: %s", err.Error())
		return nil, err
	}

	status, body, err := httpclient.PostJSON(ctx, c.URL, p, nil)
	if err != nil {
		log.Errorf(ctx, "httpclient.PostJSON error: %s", err.Error())
		return nil, err
	}
	if status != http.StatusOK {
		err = fmt.Errorf("httpclient.PostJSON status: %d", status)
		log.Errorf(ctx, "%s", err.Error())
		return nil, err
	}

	var res []*ClientResponse
	err = json.Unmarshal(body, &res)
	if err != nil {
		log.Errorf(ctx, "json.Unmarshal error: %s", err.Error())
		return nil, err
	}

	return res, nil
}

// NewClient ... Clientを作成する
func NewClient(url string, headers map[string]string) *Client {
	return &Client{
		URL:     url,
		Headers: headers,
	}
}
