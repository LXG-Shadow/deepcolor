package deepcolor

type Requester interface {
	Next(response *Response) *Request
}

type Request struct {
	Url     string            `json:"url"`
	Charset string            `json:"charset"`
	Header  map[string]string `json:"header"`
}

type SingleRequest struct {
	Url     string            `json:"url"`
	Charset string            `json:"charset"`
	Header  map[string]string `json:"header"`
	flag    bool
}

func (r *SingleRequest) Next(response *Response) *Request {
	if r.flag {
		return nil
	}
	r.flag = true
	return &Request{
		Url:     r.Url,
		Charset: r.Charset,
		Header:  r.Header,
	}
}

func NewSingleRequest(url string, charset string, header map[string]string) Requester {
	return &SingleRequest{
		Url:     url,
		Charset: charset,
		Header:  header,
		flag:    false,
	}
}
