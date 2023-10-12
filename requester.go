package deepcolor

import (
	"github.com/aynakeya/deepcolor/dphttp"
	"github.com/aynakeya/deepcolor/dphttp/requesters"
)

func init() {
	_defaultRequester = NewRestyRequester()
}

var _defaultRequester dphttp.IRequester = nil

func SetDefaultRequester(requester dphttp.IRequester) {
	_defaultRequester = requester
}

func NewRestyRequester() dphttp.IRequester {
	return requesters.NewRestyRequester()
}
