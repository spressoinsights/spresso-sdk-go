package http_client

import (
	"context"
	"time"

	"github.com/go-resty/resty/v2"
	"github.com/goccy/go-json"
)

type RestyClient interface {
	R(ctx context.Context, serviceName string, successStatusCode int) RestyRequest
}

type restyClient struct {
	underlyingClient  *resty.Client
	defaultTimeout    time.Duration
	defaultRetryCount int
}

// sets default content type to application/json
func (r *restyClient) R(ctx context.Context, serviceName string, successStatusCode int) RestyRequest {
	return &restyRequest{
		underlyingReq:     r.underlyingClient.R(),
		ctx:               ctx,
		serviceName:       serviceName,
		maxRetries:        r.defaultRetryCount,
		requestCount:      0,
		timeout:           r.defaultTimeout,
		successStatusCode: successStatusCode,
	}
}

func NewRestyClient(defaultTimeout *time.Duration, defaultRetryCount *int) RestyClient {
	client := resty.New()

	// use go-json for faster marshal/unmarshal
	client.JSONMarshal = json.Marshal
	client.JSONUnmarshal = json.Unmarshal

	timeout := 10 * time.Second
	if defaultTimeout != nil {
		timeout = *defaultTimeout
	}

	retryCount := 0
	if defaultRetryCount != nil {
		retryCount = *defaultRetryCount
	}

	return &restyClient{
		underlyingClient:  client,
		defaultTimeout:    timeout,
		defaultRetryCount: retryCount,
	}
}

type RestyRequest interface {
	SetHeader(header, value string) RestyRequest
	SetHeaderMultiValues(headers map[string][]string) RestyRequest
	SetQueryParam(param, value string) RestyRequest
	SetQueryParams(params map[string]string) RestyRequest
	SetBody(body interface{}) RestyRequest
	SetResult(res interface{}) RestyRequest
	SetPathParam(params, value string) RestyRequest
	SetPathParams(params map[string]string) RestyRequest

	SetRetryCount(retryCount int) RestyRequest
	SetTimeout(timeout time.Duration) RestyRequest

	Get(url string) (*resty.Response, error)
	Post(url string) (*resty.Response, error)
	Put(url string) (*resty.Response, error)
}

type restyRequest struct {
	underlyingReq *resty.Request

	ctx         context.Context
	serviceName string

	maxRetries   int
	requestCount int

	timeout time.Duration

	successStatusCode int
}

func (r *restyRequest) SetHeader(header, value string) RestyRequest {
	r.underlyingReq = r.underlyingReq.SetHeader(header, value)
	return r
}

func (r *restyRequest) SetHeaderMultiValues(headers map[string][]string) RestyRequest {
	r.underlyingReq = r.underlyingReq.SetHeaderMultiValues(headers)
	return r
}

func (r *restyRequest) SetQueryParam(param, value string) RestyRequest {
	r.underlyingReq = r.underlyingReq.SetQueryParam(param, value)
	return r
}

func (r *restyRequest) SetQueryParams(params map[string]string) RestyRequest {
	r.underlyingReq = r.underlyingReq.SetQueryParams(params)
	return r
}

func (r *restyRequest) SetBody(body interface{}) RestyRequest {
	r.underlyingReq = r.underlyingReq.SetBody(body)
	return r
}

func (r *restyRequest) SetResult(res interface{}) RestyRequest {
	r.underlyingReq = r.underlyingReq.SetResult(res)
	return r
}

func (r *restyRequest) SetPathParam(param, value string) RestyRequest {
	r.underlyingReq = r.underlyingReq.SetPathParam(param, value)
	return r
}

func (r *restyRequest) SetPathParams(params map[string]string) RestyRequest {
	r.underlyingReq = r.underlyingReq.SetPathParams(params)
	return r
}

func (r *restyRequest) SetRetryCount(retryCount int) RestyRequest {
	r.maxRetries = retryCount
	return r
}

func (r *restyRequest) SetTimeout(timeout time.Duration) RestyRequest {
	r.timeout = timeout
	return r
}

func (r *restyRequest) prepareRequest() (*resty.Request, context.CancelFunc) {

	req := r.underlyingReq.AddRetryCondition(func(resp *resty.Response, err error) bool {
		r.requestCount++
		return r.requestCount >= r.maxRetries
	})

	requestCtx, cancelFunc := context.WithTimeout(r.underlyingReq.Context(), r.timeout)
	return req.SetContext(requestCtx), cancelFunc
}

func (r *restyRequest) parseResult(response *resty.Response, err error) (*resty.Response, error) {
	if err == nil {

		return response, nil
	}

	return nil, nil
}

func (r *restyRequest) Get(url string) (*resty.Response, error) {
	req, cancel := r.prepareRequest()
	defer cancel()

	return r.parseResult(req.Get(url))
}

func (r *restyRequest) Post(url string) (*resty.Response, error) {
	req, cancel := r.prepareRequest()
	defer cancel()

	return r.parseResult(req.Post(url))
}

func (r *restyRequest) Put(url string) (*resty.Response, error) {
	req, cancel := r.prepareRequest()
	defer cancel()

	return r.parseResult(req.Put(url))
}
