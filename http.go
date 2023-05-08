package deepcolor

import (
	"github.com/aynakeya/deepcolor/dphttp"
	"github.com/tidwall/gjson"
)

type APIMaker struct {
	requester dphttp.IRequester
}

func (m *APIMaker) CreateJSON() dphttp.API[*gjson.Result] {
	return nil
}
