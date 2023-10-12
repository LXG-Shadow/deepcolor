package deepcolor

import "github.com/aynakeya/deepcolor/dphttp"

func CreateApiRecverFunc[P dphttp.ParameterType, T dphttp.ParserResultType, R dphttp.ResultType](
	request dphttp.RequestFunc[P], parser dphttp.ParserFunc[T], result dphttp.ResultFunc[T, R]) dphttp.ApiRecverFunc[P, R] {
	return dphttp.NewRecverFunc(_defaultRequester, &dphttp.ApiInfo[P, T, R]{
		Request: request,
		Parser:  parser,
		Result:  result,
	})
}

//func CreateChainApiFunc[P dphttp.ParameterType, T dphttp.ParserResultType, R dphttp.ResultType](
//	requester dphttp.IRequester,
//	request dphttp.RequestFunc[P], parser dphttp.ParserFunc[T], result dphttp.ResultFunc[T, R],
//	next dphttp.NextRequestFunc[P, T, R]) dphttp.ApiRecverFunc[P, R] {
//	return dphttp.CreateRecverFunc(requester, &dphttp.ApiInfo[P, T, R]{
//		Request: request,
//		Parser:  parser,
//		Result:  result,
//		Next:    next,
//	})
//}

func CreateApiResultFunc[P dphttp.ParameterType, T dphttp.ParserResultType, R dphttp.ResultType](
	request dphttp.RequestFunc[P], parser dphttp.ParserFunc[T], result dphttp.ResultFunc[T, *R]) dphttp.ApiResultFunc[P, R] {
	return dphttp.NewResultFunc(_defaultRequester, &dphttp.ApiInfo[P, T, *R]{
		Request: request,
		Parser:  parser,
		Result:  result,
	})
}
