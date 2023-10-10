package deepcolor

import (
	"fmt"
	"github.com/aynakeya/deepcolor/dphttp"
)

func NewGetRequestWithSingleQuery(
	uri string,
	query string, headers map[string]string) dphttp.RequestFunc[string] {
	return func(param string) (*dphttp.Request, error) {
		url := dphttp.UrlMustParse(uri)
		paramVals := url.Query()
		paramVals.Set(query, param)
		url.RawQuery = paramVals.Encode()
		return &dphttp.Request{
			Method: dphttp.GET,
			Url:    url,
			Header: headers,
		}, nil
	}
}

func NewGetRequestWithQuery(
	uri string,
	queries []string, headers map[string]string) dphttp.RequestFunc[[]string] {
	return func(params []string) (*dphttp.Request, error) {
		if len(queries) > len(params) {
			return nil, fmt.Errorf("only receive %d parameter, required %d", len(params), len(queries))
		}
		url := dphttp.UrlMustParse(uri)
		paramVals := url.Query()
		for i, _ := range queries {
			paramVals.Set(queries[i], params[i])
		}
		url.RawQuery = paramVals.Encode()
		return &dphttp.Request{
			Method: dphttp.GET,
			Url:    url,
			Header: headers,
		}, nil
	}
}
