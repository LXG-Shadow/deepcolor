package deepcolor

import (
	"net/http"
	"strings"
)

type ResultParser struct {
	Type      ResultType
	Data      interface{}
	GetValue  func(i interface{}, rule ItemRule) string
	GetValues func(i interface{}, rule ItemRule) []string
}

// Response struct
// adapt from resty response.go
type Response struct {
	RawResponse *http.Response
	body        []byte
	size        int64
}

func (r *Response) Body() []byte {
	if r.RawResponse == nil {
		return []byte{}
	}
	return r.body
}

func (r *Response) StatusCode() int {
	if r.RawResponse == nil {
		return 0
	}
	return r.RawResponse.StatusCode
}

func (r *Response) Header() http.Header {
	if r.RawResponse == nil {
		return http.Header{}
	}
	return r.RawResponse.Header
}

func (r *Response) String() string {
	if r.body == nil {
		return ""
	}
	return strings.TrimSpace(string(r.body))
}

//
//func (t Response) GetSingle(item Item) (result string) {
//	result = ""
//	if item.Type != ItemTypeSingle {
//		return
//	}
//	if len(item.Rules) < 1 {
//		return
//	}
//	var parser *ResultParser
//	for _, rule := range item.Rules {
//		if parser = t.GetParser(rule.Selector.Type.GetValidResultType()); parser != nil {
//			result += parser.GetValue(parser.Data, rule)
//		}
//	}
//	return
//}
//
//func (t Response) GetList(item Item) (result []string) {
//	result = make([]string, 0)
//	if item.Type != ItemTypeList {
//		return
//	}
//	if len(item.Rules) < 1 {
//		return
//	}
//	var parser *ResultParser
//	for _, rule := range item.Rules {
//		parser = t.GetParser(rule.Selector.Type.GetValidResultType())
//		if parser == nil {
//			continue
//		}
//		for i, val := range parser.GetValues(parser.Data, rule) {
//			if len(result) <= i {
//				result = append(result, val)
//			} else {
//				result[i] += val
//			}
//		}
//	}
//	return
//}
//
//func (t Response) GetMap(item Item) (result map[string]string) {
//	result = map[string]string{}
//	if item.Type != ItemTypeMap {
//		return
//	}
//	var parser *ResultParser
//	for _, rule := range item.Rules {
//		parser = t.GetParser(rule.Selector.Type.GetValidResultType())
//		if parser == nil {
//			continue
//		}
//		result[rule.Key] += parser.GetValue(parser.Data, rule)
//	}
//	return
//}
//
//func (t Response) GetMapList(item Item) (result []map[string]string) {
//	result = make([]map[string]string, 0)
//	if item.Type != ItemTypeMapList {
//		return
//	}
//	var parser *ResultParser
//	for _, rule := range item.Rules {
//		parser = t.GetParser(rule.Selector.Type.GetValidResultType())
//		if parser == nil {
//			continue
//		}
//		for i, val := range parser.GetValues(parser.Data, rule) {
//			if len(result) <= i {
//				result = append(result, newBaseMap(item))
//			}
//			result[i][rule.Key] += val
//		}
//	}
//	return
//}
