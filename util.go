package deepcolor

import (
	"github.com/aynakeya/deepcolor/dphttp"
	"net/url"
)

func ParseUrl(rawurl string) *url.URL {
	return dphttp.UrlMustParse(rawurl)
}
