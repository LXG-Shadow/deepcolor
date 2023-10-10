package deepcolor

import "github.com/aynakeya/deepcolor/dphttp"

func CreateApiFunc[P dphttp.ParameterType, T dphttp.ParserResultType, R dphttp.ResultType](
	requester dphttp.IRequester,
	request dphttp.RequestFunc[P], parser dphttp.ParserFunc[T], result dphttp.ResultFunc[T, R]) dphttp.ApiFunc[P, R] {
	return dphttp.CreateApiFunc(requester, &dphttp.ApiInfo[P, T, R]{
		Request: request,
		Parser:  parser,
		Result:  result,
	})
}

func CreateChainApiFunc[P dphttp.ParameterType, T dphttp.ParserResultType, R dphttp.ResultType](
	requester dphttp.IRequester,
	request dphttp.RequestFunc[P], parser dphttp.ParserFunc[T], result dphttp.ResultFunc[T, R],
	next dphttp.NextRequestFunc[P, T, R]) dphttp.ApiFunc[P, R] {
	return dphttp.CreateApiFunc(requester, &dphttp.ApiInfo[P, T, R]{
		Request: request,
		Parser:  parser,
		Result:  result,
		Next:    next,
	})
}

func CreateApiResultFunc[P dphttp.ParameterType, T dphttp.ParserResultType, R dphttp.ResultType](
	requester dphttp.IRequester,
	request dphttp.RequestFunc[P], parser dphttp.ParserFunc[T], result dphttp.ResultFunc[T, *R]) dphttp.ApiFuncResult[P, R] {
	return dphttp.CreateResultFunc(requester, &dphttp.ApiInfo[P, T, *R]{
		Request: request,
		Parser:  parser,
		Result:  result,
	})
}
