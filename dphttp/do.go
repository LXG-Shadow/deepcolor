package dphttp

func FetchParsedResult[P ParserResultType](requester IRequester, request *Request, parserFunc ParserFunc[P]) (P, error) {
	httpResp, err := requester.HTTP(request)
	if err != nil {
		return *new(P), err
	}
	return parserFunc(httpResp)
}
