package deepcolor

import (
	"github.com/tidwall/gjson"
)

type RequestFunc func(uri string, header map[string]string) *HttpResponse
type RequestHandler func(tentacle *Tentacle) bool
type ResponseHandler func(result *TentacleResult) bool

func postprocess(tentacleResult *TentacleResult, handlers []ResponseHandler) bool {
	for _, handler := range handlers {
		if !handler(tentacleResult) {
			return false
		}
	}
	return true
}

func preprocess(tentacle *Tentacle, handlers []RequestHandler) bool {
	for _, handler := range handlers {
		if !handler(tentacle) {
			return false
		}
	}
	return true
}

func Fetch(tentacle Tentacle, requestFunc RequestFunc,
	preHandlers []RequestHandler, postHandlers []ResponseHandler) (result *TentacleResult, err error) {
	if !preprocess(&tentacle, preHandlers) {
		return nil, ErrorRequestCancelByPreprocessFunction
	}
	httpResult := requestFunc(tentacle.Url, tentacle.Header).String()
	if httpResult == "" {
		return nil, ErrorHttpConnectionFail
	}
	result = &TentacleResult{
		Request: tentacle,
		Parsers: [3]*ResultParser{},
	}
	err = MakeParser(result, httpResult)

	if err == nil {
		defer postprocess(result, postHandlers)
	}
	return result, err
}

func MakeParser(result *TentacleResult, httpResult string) error {
	r_type := result.Request.ContentType
	if r_type.Contains(ResultTypeText) {
		//fmt.Println("initialize text parser")
		result.Parsers[0] = &ResultParser{
			Type:      ResultTypeText,
			Data:      httpResult,
			GetValue:  getTextValue,
			GetValues: getTextValues,
		}
	}
	if r_type.Contains(ResultTypeHTMl) {
		//fmt.Println("initialize html parser")
		doc, err := NewDocumentFromStringWithEncoding(httpResult, result.Request.Charset)
		if err != nil {
			return err
		}
		result.Parsers[1] = &ResultParser{
			Type:      ResultTypeHTMl,
			Data:      doc,
			GetValue:  getHTMLValue,
			GetValues: getHTMLValues,
		}
	}
	if r_type.Contains(ResultTypeJson) {
		g := gjson.Parse(httpResult)
		result.Parsers[2] = &ResultParser{
			Type:      ResultTypeJson,
			Data:      g,
			GetValue:  getJSONValue,
			GetValues: getJSONValues,
		}
	}
	return nil
}
