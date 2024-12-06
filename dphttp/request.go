package dphttp

import (
	"context"
	"net/url"
)

type Request struct {
	Method  string            `json:"method"`
	Url     *url.URL          `json:"url"`
	Header  map[string]string `json:"header"`
	Data    []byte            `json:"data"`
	Timeout int               `json:"timeout"`
	Context context.Context   `json:"-"`
}
