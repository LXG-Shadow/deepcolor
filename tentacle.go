package deepcolor

type TentacleContentType string

const (
	TentacleContentTypeText = "text"
	TentacleContentTypeHTMl = "html"
	TentacleContentTypeJson = "json"
)

type Tentacle struct {
	Url         string              `json:"url"`
	Charset     string              `json:"charset"`
	ContentType TentacleContentType `json:"content_type"`
	Header      map[string]string   `json:"header"`
}

func TentacleHTML(uri, charset string) Tentacle {
	return Tentacle{
		Url:         uri,
		Charset:     charset,
		ContentType: TentacleContentTypeHTMl,
	}
}

type TentacleResult interface {
	GetRequest() *Tentacle
	GetSingle(item Item) string
	GetList(item Item) []string
	GetMap(item Item) map[string]string
	GetMapList(item Item) []map[string]string
}

type ResultParser struct {
	GetValue  func(i interface{}, rule ItemRule) string
	GetValues func(i interface{}, rule ItemRule) []string
}

var (
	HTMLResultParser = &ResultParser{
		getHTMLValue,
		getHTMLValues,
	}
	TextResultParser = &ResultParser{
		getTextValue,
		getTextValues,
	}
)

type TentacleTextResult struct {
	TentacleResult
	Request *Tentacle
	Parser  *ResultParser
	Data    interface{}
}

func (t TentacleTextResult) GetRequest() *Tentacle{
	return t.Request
}

func (t TentacleTextResult) GetSingle(item Item) (result string) {
	result = ""
	if item.Type != ItemTypeSingle {
		return
	}
	if len(item.Rules) < 1 {
		return
	}
	for _, rule := range item.Rules {
		result += t.Parser.GetValue(t, rule)
	}
	return
}

func (t TentacleTextResult) GetList(item Item) (result []string) {
	result = make([]string, 0)
	if item.Type != ItemTypeList {
		return
	}
	if len(item.Rules) < 1 {
		return
	}
	for _, rule := range item.Rules {
		for i, val := range t.Parser.GetValues(t, rule) {
			if len(result) <= i {
				result = append(result, val)
			} else {
				result[i] += val
			}
		}
	}
	return
}

func (t TentacleTextResult) GetMap(item Item) (result map[string]string) {
	result = map[string]string{}
	if item.Type != ItemTypeMap {
		return
	}
	for _, rule := range item.Rules {
		result[rule.Key] += t.Parser.GetValue(t, rule)
	}
	return
}

func (t TentacleTextResult) GetMapList(item Item) (result []map[string]string) {
	result = make([]map[string]string, 0)
	if item.Type != ItemTypeMapList {
		return
	}
	for _, rule := range item.Rules {
		for i, val := range t.Parser.GetValues(t, rule) {
			if len(result) <= i {
				result = append(result, newBaseMap(item))
			}
			result[i][rule.Key] += val
		}
	}
	return
}

func NewTentacleWithParser(request Tentacle, data interface{}, parser *ResultParser) TentacleResult {
	return TentacleTextResult{
		Request: &request,
		Parser:  parser,
		Data:    data,
	}
}

//func (t TentacleHTMLResult) GetSingle(item Item) (result string) {
//	result = ""
//	if item.Type != ItemTypeSingle {
//		return
//	}
//	if len(item.Rules) < 1 {
//		return
//	}
//	for _, rule := range item.Rules {
//		result += getHTMLValue(&t,rule)
//	}
//	return
//}
//
//func (t TentacleHTMLResult) GetList(item Item) (result []string) {
//	result = make([]string, 0)
//	if item.Type != ItemTypeList {
//		return
//	}
//	if len(item.Rules) < 1 {
//		return
//	}
//	for _, rule := range item.Rules {
//		for i,val := range getHTMLValues(&t,rule){
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
//func (t TentacleHTMLResult) GetMap(item Item) (result map[string]string) {
//	result = map[string]string{}
//	if item.Type != ItemTypeMap {
//		return
//	}
//	for _, rule := range item.Rules {
//		result[rule.Key] += getHTMLValue(&t,rule)
//	}
//	return
//}
//
//// Deprecated: use different list instead
//func (t TentacleHTMLResult) GetMapList(item Item) (result []map[string]string) {
//	result = make([]map[string]string, 0)
//	if item.Type != ItemTypeMapList {
//		return
//	}
//	for _, rule := range item.Rules {
//		for i,val := range getHTMLValues(&t,rule){
//			if len(result) <= i {
//				result = append(result, newBaseMap(item))
//			}
//			result[i][rule.Key] += val
//		}
//	}
//	return
//}
