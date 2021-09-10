package deepcolor

import (
	"net/url"
	"regexp"
)

func IsUrl(url string) bool {
	urlRegExp := regexp.MustCompile(
		"(?i)^(?:http|ftp)s?://" +
			"(?:(?:[A-Z0-9](?:[A-Z0-9-]{0,61}[A-Z0-9])?\\.)+(?:[A-Z]{2,6}\\.?|[A-Z0-9-]{2,}\\.?)|" +
			"localhost|" +
			"\\d{1,3}\\.\\d{1,3}\\.\\d{1,3}\\.\\d{1,3})" +
			"(?::\\d+)?" +
			"(?:/?|[/?]\\S+)$")
	return urlRegExp.FindString(url) != ""
}

func QueryEscapeWithEncoding(str string, encoding string) string {
	return url.QueryEscape(EncodeString(str, encoding))
}

func GetUrlHost(uri string) string {
	u, err := url.Parse(uri)
	if err != nil {
		return ""
	}
	return u.Host
}