package deepcolor

type ResponseParser interface {
	Initialize(resp *Response) error
	Get(rule *Selector) interface{}
	GetValue(rule *Selector) interface{}
	GetValues(rule *Selector) []interface{}
}
