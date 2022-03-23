package deepcolor

type RequestFunc func(uri string, header map[string]string) *HttpResponse
type RequestHandler func(tentacle Tentacle) bool
type ResponseHandler func(result TentacleResult) bool

func postprocess(tentacleResult TentacleResult, handlers []ResponseHandler) bool {
	for _, handler := range handlers {
		if !handler(tentacleResult) {
			return false
		}
	}
	return true
}

func preprocess(tentacle Tentacle, handlers []RequestHandler) bool {
	for _, handler := range handlers {
		if !handler(tentacle) {
			return false
		}
	}
	return true
}

func Fetch(tentacle Tentacle, requestFunc RequestFunc,
	preHandlers []RequestHandler, postHandlers []ResponseHandler) (result TentacleResult, err error) {
	switch tentacle.ContentType {
	case TentacleContentTypeHTMl:
		return FetchHTML(tentacle, requestFunc, preHandlers, postHandlers)
	case TentacleContentTypeText:
		return FetchText(tentacle, requestFunc, preHandlers, postHandlers)
	default:
		return FetchHTML(tentacle, requestFunc, preHandlers, postHandlers)
	}
}

func FetchText(tentacle Tentacle, requestFunc RequestFunc,
	preHandlers []RequestHandler, postHandlers []ResponseHandler) (result TentacleResult, err error) {
	if !preprocess(tentacle, preHandlers) {
		return nil, ErrorRequestCancelByPreprocessFunction
	}
	httpResult := requestFunc(tentacle.Url, tentacle.Header).String()
	if httpResult == "" {
		return nil, ErrorHttpConnectionFail
	}
	tentacleResult := NewTentacleWithParser(tentacle, httpResult, TextResultParser)
	defer postprocess(tentacleResult, postHandlers)
	return tentacleResult, nil
}

func FetchHTML(tentacle Tentacle, requestFunc RequestFunc,
	preHandlers []RequestHandler, postHandlers []ResponseHandler) (result TentacleResult, err error) {
	if !preprocess(tentacle, preHandlers) {
		return nil, ErrorRequestCancelByPreprocessFunction
	}
	httpResult := requestFunc(tentacle.Url, tentacle.Header).String()
	if httpResult == "" {
		return nil, ErrorHttpConnectionFail
	}
	doc, err := NewDocumentFromStringWithEncoding(httpResult, tentacle.Charset)
	if err != nil {
		return nil, err
	}
	tentacleResult := NewTentacleWithParser(tentacle, doc, HTMLResultParser)
	defer postprocess(tentacleResult, postHandlers)
	return tentacleResult, nil
}
