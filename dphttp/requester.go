package dphttp

import "net/http"

type Config struct {
	BaseUrl string
	Header  map[string]string
	Cookie  map[string]string
	Timeout int
}

func NewConfig() *Config {
	return &Config{
		BaseUrl: "",
		Header:  make(map[string]string),
		Cookie:  make(map[string]string),
		Timeout: 10,
	}
}

type IBaseRequester interface {
	// Config return a modifiable Config struct
	Config() *Config
	HTTP(req *Request) (*Response, error)
}

// IRequester is the interface for all requesters
type IRequester interface {
	IBaseRequester
	Get(uri string, headers map[string]string) (*Response, error)
	Post(uri string, headers map[string]string, body any) (*Response, error)
	GetQuery(uri string, query map[string]string, headers map[string]string) (*Response, error)
	GetX(uri string) (*Response, error)
	PostX(uri string, body any) (*Response, error)
	GetQueryX(uri string, query map[string]string) (*Response, error)
}

func NewRequester(base IBaseRequester) IRequester {
	return &requester{
		base: base,
	}
}

type requester struct {
	base IBaseRequester
}

// Config return a modifiable Config struct
func (r *requester) Config() *Config {
	return r.base.Config()
}

func (r *requester) HTTP(req *Request) (*Response, error) {
	return r.base.HTTP(req)
}

func (r *requester) Get(uri string, headers map[string]string) (*Response, error) {
	return r.base.HTTP(&Request{
		Method: http.MethodGet,
		Url:    BuildUrl(r.Config().BaseUrl, uri),
		Header: headers,
	})
}

func (r *requester) Post(uri string, headers map[string]string, body any) (*Response, error) {
	return r.base.HTTP(&Request{
		Method: http.MethodPost,
		Url:    BuildUrl(r.Config().BaseUrl, uri),
		Header: headers,
		Data:   FormatBodyData(body),
	})
}

func (r *requester) GetQuery(uri string, query map[string]string, headers map[string]string) (*Response, error) {
	u := BuildUrl(r.Config().BaseUrl, uri)
	paramVals := u.Query()
	for key, value := range query {
		paramVals.Set(key, value)
	}
	u.RawQuery = paramVals.Encode()
	return r.base.HTTP(&Request{
		Method: http.MethodGet,
		Url:    u,
		Header: headers,
	})
}

func (r *requester) GetX(uri string) (*Response, error) {
	return r.base.HTTP(&Request{
		Method: http.MethodGet,
		Url:    BuildUrl(r.Config().BaseUrl, uri),
	})
}

func (r *requester) PostX(uri string, body any) (*Response, error) {
	return r.base.HTTP(&Request{
		Method: http.MethodPost,
		Url:    BuildUrl(r.Config().BaseUrl, uri),
		Data:   FormatBodyData(body),
	})
}

func (r *requester) GetQueryX(uri string, query map[string]string) (*Response, error) {
	u := BuildUrl(r.Config().BaseUrl, uri)
	paramVals := u.Query()
	for key, value := range query {
		paramVals.Set(key, value)
	}
	u.RawQuery = paramVals.Encode()
	return r.base.HTTP(&Request{
		Method: http.MethodGet,
		Url:    u,
	})
}
