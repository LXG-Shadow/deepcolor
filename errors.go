package deepcolor

import "errors"

var (
	ErrorRequestCancelByPreprocessFunction = errors.New("request cancel by preprocess function")
	ErrorHttpConnectionFail                = errors.New("http connection error")
	ErrorEmptyBody                         = errors.New("response body is empty")
)
