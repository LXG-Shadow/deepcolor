package dprequest

import (
	"encoding/json"
	"net/url"
	"regexp"
)

var absoluteUrlRegex = regexp.MustCompile(`([a-z][a-z\d+\-.]*:)?//`)
var baseUrlRegex = regexp.MustCompile(`/+$`)
var refUrlRegex = regexp.MustCompile(`^/+`)

func FormatBodyData(data any) []byte {
	switch data.(type) {
	case string:
		return []byte(data.(string))
	case []byte:
		return data.([]byte)
	}
	rs, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}
	return rs
}

func headerMerge(src map[string]string, updated map[string]string) map[string]string {
	header := make(map[string]string)
	for k, v := range src {
		header[k] = v
	}
	for k, v := range updated {
		header[k] = v
	}
	return header
}

func UrlMustParse(rawurl string) *url.URL {
	u, err := url.Parse(rawurl)
	if err != nil {
		panic(err)
	}
	return u
}

// UrlJoin adapted from axios core
func BuildUrl(base string, ref string) *url.URL {
	if base == "" || absoluteUrlRegex.MatchString(ref) {
		return UrlMustParse(ref)
	}
	// ????
	if ref == "" {
		return UrlMustParse(ref)
	}
	return UrlMustParse(baseUrlRegex.ReplaceAllString(base, "") + "/" + refUrlRegex.ReplaceAllString(ref, ""))
}
