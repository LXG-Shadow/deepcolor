package dprequest

import (
	"fmt"
	"github.com/go-resty/resty/v2"
	"golang.org/x/net/context"
	"time"
)

type restyRequester struct {
	config *Config
	client *resty.Client
}

func NewRestyRequester() IRequester {
	return NewRequester(&restyRequester{
		client: resty.New(),
		config: &Config{Timeout: 10},
	})
}

func (r *restyRequester) Config() *Config {
	return r.config
}

func (r *restyRequester) HTTP(req *Request) (*Response, error) {
	var resp *resty.Response
	var err error
	req.Header = headerMerge(req.Header, r.config.Header)
	if req.Timeout == 0 {
		req.Timeout = r.config.Timeout
	}
	if req.Context == nil {
		req.Context = context.Background()
	}
	ctx, _ := context.WithTimeout(req.Context, time.Duration(req.Timeout)*time.Second)
	switch req.Method {
	case GET:
		resp, err = r.client.R().
			SetContext(ctx).
			SetHeaders(req.Header).
			Get(req.Url.String())
	case POST:
		fmt.Println("POST", req.Url.String(), req.Header, req.Data)
		resp, err = r.client.R().
			SetContext(ctx).
			SetHeaders(req.Header).
			SetBody(req.Data).
			Post(req.Url.String())
	case HEAD:
		resp, err = r.client.R().
			SetContext(ctx).
			SetHeaders(req.Header).
			Head(req.Url.String())
	}
	if err != nil {
		return &Response{
			Request: req,
		}, err
	}
	return &Response{
		Request:     req,
		RawResponse: resp.RawResponse,
		body:        resp.Body(),
		size:        resp.Size(),
	}, nil
}
