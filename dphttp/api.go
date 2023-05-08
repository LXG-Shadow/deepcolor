package dphttp

type ParserResultType any
type ParserFunc[P ParserResultType] func(response *Response) (P, error)
type ResultType any
type ResultFunc[P ParserResultType, R ResultType] func(result P) (R, error)

type ApiInfo[P ParserResultType, R ResultType] struct {
	Request Request
	Parser  ParserFunc[P]
	Handler ResultFunc[P, R]
}

type API[R ResultType] func() (R, error)

func CreateAPI[P ParserResultType, R ResultType](requester IRequester, api *ApiInfo[P, R]) API[R] {
	return func() (R, error) {
		httpResp, err := requester.HTTP(&api.Request)
		if err != nil {
			return *new(R), err
		}
		parsedResult, err := api.Parser(httpResp)
		if err != nil {
			return *new(R), err
		}
		return api.Handler(parsedResult)
	}
}
