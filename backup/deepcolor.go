package backup

import (
	"github.com/aynakeya/deepcolor/transform"
)

type RequestFunc func(req *Request) *Response
type RequestHandler func(request *Request) bool
type ResponseHandler func(response *Response) bool
type TentacleHandler func(tentacle *Tentacle)

type Deepcolor struct {
	ReqFunc     RequestFunc
	Requester   Requester
	ReqHandler  []RequestHandler
	RespHandler []ResponseHandler
	Tentacles   []*Tentacle
}

func (d *Deepcolor) OnRequest(handlers ...RequestHandler) {
	d.ReqHandler = append(d.ReqHandler, handlers...)
}

func (d *Deepcolor) OnResponse(handlers ...ResponseHandler) {
	d.RespHandler = append(d.RespHandler, handlers...)
}

type TentacleMapper struct {
	Selector   *Selector
	Translator transform.Translator
}

func (s *deepcolor.Selector) ToMapper() *TentacleMapper {
	return &TentacleMapper{
		Selector:   s,
		Translator: nil,
	}
}

type Tentacle struct {
	Parser       ResponseParser
	ValueMapper  map[string]*TentacleMapper
	Transformers []*transform.Transformer
	Handlers     []TentacleHandler
}

func (t *Tentacle) Initialize(response *Response) error {
	return t.Parser.Initialize(response)
}

func (t *Tentacle) GetItems() map[string]interface{} {
	items := make(map[string]interface{})
	for key, rule := range t.ValueMapper {
		v := t.Parser.Get(rule.Selector)
		if rule.Translator != nil {
			v = rule.Translator.MustApply(v)
		}
		items[key] = v
	}
	return items
}

func (t *Tentacle) Extract(value interface{}) {
	for key, rule := range t.ValueMapper {
		if v, ok := transform.Field(key).GetValueE(value); ok {
			fv := t.Parser.Get(rule.Selector)
			if rule.Translator != nil {
				fv = rule.Translator.MustApply(fv)
			}
			transform.SetFieldValue(fv, v)
		}
	}
}

func (t *Tentacle) Transform(value interface{}) error {
	for _, tran := range t.Transformers {
		err := tran.Transform(value)
		if err != nil {
			return err
		}
	}
	return nil
}

func (t *Tentacle) ExtractAndTransform(value interface{}) error {
	t.Extract(value)
	return t.Transform(value)
}
