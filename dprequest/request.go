package dprequest

import (
	"context"
	"net/url"
)

type RequestMethod string

const (
	GET  RequestMethod = "GET"
	POST RequestMethod = "POST"
	HEAD RequestMethod = "HEAD"
)

type Request struct {
	Method  RequestMethod     `json:"method"`
	Url     *url.URL          `json:"url"`
	Header  map[string]string `json:"header"`
	Data    []byte            `json:"data"`
	Timeout int               `json:"timeout"`
	Context context.Context   `json:"-"`
}
