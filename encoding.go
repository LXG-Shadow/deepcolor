package deepcolor

import (
	"golang.org/x/net/html/charset"
	"golang.org/x/text/transform"
	"io/ioutil"
	"log"
	"strings"
)

func transformString(t transform.Transformer, s string) (string, error) {
	r := transform.NewReader(strings.NewReader(s), t)
	b, err := ioutil.ReadAll(r)
	return string(b), err
}

func DecodeString(str string, encoding string) string {
	e, _ := charset.Lookup(encoding)
	if e == nil {
		log.Println("encoding not found: ", encoding)
		return ""
	}
	decodeStr, err := transformString(e.NewDecoder(), str)
	if err != nil {
		log.Println("error", encoding)
		return ""
	}
	return decodeStr
}

func EncodeString(str, encoding string) string {
	e, _ := charset.Lookup(encoding)
	if e == nil {
		log.Println("encoding not found: ", encoding)
		return ""
	}
	encodeStr, err := transformString(e.NewEncoder(), str)
	if err != nil {
		log.Println("error", encoding)
		return ""
	}
	return encodeStr
}
