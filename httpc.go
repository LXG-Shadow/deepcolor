package deepcolor

import (
	"github.com/go-resty/resty/v2"
	"math/rand"
	"time"
)

var fakeUserAgents = []string{
	"Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/60.0.3112.101 Safari/537.36",
	"Mozilla/5.0 (compatible; MSIE 9.0; Windows NT 6.1; WOW64; Trident/5.0; SLCC2; Media Center PC 6.0; InfoPath.3; MS-RTC LM 8; Zune 4.7)",
	"Mozilla/4.0 (compatible; MSIE 9.0; Windows NT 5.1; Trident/5.0)",
	"Mozilla/5.0 (X11; Linux x86_64; rv:2.2a1pre) Gecko/20100101 Firefox/4.2a1pre",
	"Mozilla/5.0 (Windows NT 6.1; Win64; x64; rv:2.0b11pre) Gecko/20110131 Firefox/4.0b11pre",
	"Mozilla/5.0 (X11; U; Linux i686; ru-RU; rv:1.9.2a1pre) Gecko/20090405 Ubuntu/9.04 (jaunty) Firefox/3.6a1pre",
	"Mozilla/5.0 (X11; U; Linux x86_64; en-US; rv:1.9.2.8) Gecko/20100723 SUSE/3.6.8-0.1.1 Firefox/3.6.8",
	"Mozilla/5.0 (Windows; U; Windows NT 6.1; pt-PT; rv:1.9.2.6) Gecko/20100625 Firefox/3.6.6",
	"Mozilla/5.0 (Windows; U; Windows NT 6.1; it; rv:1.9.2.6) Gecko/20100625 Firefox/3.6.6 ( .NET CLR 3.5.30729)",
	"Mozilla/5.0 (Windows; U; Windows NT 6.1; en-US; rv:1.9.2.6) Gecko/20100625 Firefox/3.6.6 (.NET CLR 3.5.30729)",
	"Mozilla/5.0 (Windows; U; Windows NT 6.1; ru; rv:1.9.2.4) Gecko/20100513 Firefox/3.6.4",
	"Mozilla/5.0 (Windows; U; Windows NT 6.1; ja; rv:1.9.2.4) Gecko/20100611 Firefox/3.6.4 GTB7.1",
}

func updateHeader(headers ...map[string]string) map[string]string {
	header := map[string]string{}
	for _, h := range headers {
		for key, val := range h {
			header[key] = val
		}
	}
	return header
}

func GetRandomUserAgent() string {
	random := rand.New(rand.NewSource(time.Now().UnixNano()))
	return fakeUserAgents[random.Intn(len(fakeUserAgents))]
}

func Get(url string, header map[string]string) *Response {
	resp, err := resty.New().R().
		SetHeaders(header).
		Get(url)
	if err != nil {
		return nil
	}
	return &Response{
		RawResponse: resp.RawResponse,
		body:        resp.Body(),
		size:        resp.Size(),
	}

}

func GetCORS(uri string, header map[string]string) *Response {
	host := GetUrlHost(uri)
	if header == nil {
		header = map[string]string{}
	}
	header = updateHeader(header, map[string]string{
		"origin":     host,
		"referer":    host,
		"user-agent": GetRandomUserAgent()})
	return Get(uri, header)
}

func Post(url string, header map[string]string, data map[string]string) *Response {
	resp, err := resty.New().R().
		SetHeaders(header).
		SetFormData(data).
		Post(url)
	if err != nil {
		return nil
	}
	return &Response{
		RawResponse: resp.RawResponse,
		body:        resp.Body(),
		size:        resp.Size(),
	}
}
