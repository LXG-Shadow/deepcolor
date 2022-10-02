package deepcolor

import (
	"net/http"
	"strings"
)

// Response struct
// adapt from resty response.go
type Response struct {
	Request     *Request
	RawResponse *http.Response
	body        []byte
	size        int64
}

func (r *Response) Body() []byte {
	if r.RawResponse == nil {
		return []byte{}
	}
	return r.body
}

func (r *Response) StatusCode() int {
	if r.RawResponse == nil {
		return 0
	}
	return r.RawResponse.StatusCode
}

func (r *Response) Header() http.Header {
	if r.RawResponse == nil {
		return http.Header{}
	}
	return r.RawResponse.Header
}

func (r *Response) String() string {
	if r.body == nil {
		return ""
	}
	return strings.TrimSpace(string(r.body))
}
