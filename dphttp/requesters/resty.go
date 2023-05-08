package requesters

import (
	"github.com/aynakeya/deepcolor/dphttp"
	"github.com/aynakeya/deepcolor/pkg/dputil"
	"github.com/go-resty/resty/v2"
	"golang.org/x/net/context"
	"time"
)

type restyRequester struct {
	config *dphttp.Config
	client *resty.Client
}

func NewRestyRequester() dphttp.IRequester {
	return dphttp.NewRequester(&restyRequester{
		client: resty.New(),
		config: &dphttp.Config{Timeout: 10},
	})
}

func (r *restyRequester) Config() *dphttp.Config {
	return r.config
}

func (r *restyRequester) HTTP(req *dphttp.Request) (*dphttp.Response, error) {
	var resp *resty.Response
	var err error
	req.Header = dputil.HttpHeaderMerge(req.Header, r.config.Header)
	if req.Timeout == 0 {
		req.Timeout = r.config.Timeout
	}
	if req.Context == nil {
		req.Context = context.Background()
	}
	ctx, _ := context.WithTimeout(req.Context, time.Duration(req.Timeout)*time.Second)
	switch req.Method {
	case dphttp.GET:
		resp, err = r.client.R().
			SetContext(ctx).
			SetHeaders(req.Header).
			Get(req.Url.String())
	case dphttp.POST:
		resp, err = r.client.R().
			SetContext(ctx).
			SetHeaders(req.Header).
			SetBody(req.Data).
			Post(req.Url.String())
	case dphttp.HEAD:
		resp, err = r.client.R().
			SetContext(ctx).
			SetHeaders(req.Header).
			Head(req.Url.String())
	}
	if err != nil {
		return &dphttp.Response{
			Request: req,
		}, err
	}
	return &dphttp.Response{
		Request:     req,
		RawResponse: resp.RawResponse,
		RawBody:     resp.Body(),
		Size:        resp.Size(),
	}, nil
}
