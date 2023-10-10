package deepcolor

import (
	"github.com/aynakeya/deepcolor/dphttp"
	"github.com/aynakeya/deepcolor/dphttp/requesters"
)

func NewRestyRequester() dphttp.IRequester {
	return requesters.NewRestyRequester()
}
