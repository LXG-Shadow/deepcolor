package deepcolor

type ResultType int

func (t ResultType) Contains(resultType ResultType) bool {
	return (t & resultType) != 0
}

const (
	ResultTypeText ResultType = 0b1
	ResultTypeHTMl ResultType = 0b10
	ResultTypeJson ResultType = 0b100
	ResultTypeRSS  ResultType = 0b1000
)

type Requester interface {
	Next(response *Response) (*Request, bool)
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
}

func (r *SingleRequest) Next(response *Response) (*Request, bool) {
	return &Request{
		Url:     r.Url,
		Charset: r.Charset,
		Header:  r.Header,
	}, false
}
