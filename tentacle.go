package deepcolor

type Tentacle struct {
	Url         string            `json:"url"`
	Charset     string            `json:"charset"`
	ContentType ResultType        `json:"content_type"`
	Header      map[string]string `json:"header"`
}

func TentacleHTML(uri, charset string) Tentacle {
	return Tentacle{
		Url:         uri,
		Charset:     charset,
		ContentType: ResultTypeHTMl | ResultTypeText,
	}
}

func TentacleJson(uri, charset string) Tentacle {
	return Tentacle{
		Url:         uri,
		Charset:     charset,
		ContentType: ResultTypeJson | ResultTypeText,
	}
}

type ResultParser struct {
	Type      ResultType
	Data      interface{}
	GetValue  func(i interface{}, rule ItemRule) string
	GetValues func(i interface{}, rule ItemRule) []string
}

type TentacleResult struct {
	Request Tentacle
	Parsers [3]*ResultParser
}

func (t TentacleResult) GetParser(content_type ResultType) *ResultParser {
	switch {
	case (content_type & ResultTypeText) != 0:
		return t.Parsers[0]
	case (content_type & ResultTypeHTMl) != 0:
		return t.Parsers[1]
	case (content_type & ResultTypeJson) != 0:
		return t.Parsers[2]
	default:
		return nil
	}
}

func (t TentacleResult) GetRequest() Tentacle {
	return t.Request
}

func (t TentacleResult) GetSingle(item Item) (result string) {
	result = ""
	if item.Type != ItemTypeSingle {
		return
	}
	if len(item.Rules) < 1 {
		return
	}
	var parser *ResultParser
	for _, rule := range item.Rules {
		if parser = t.GetParser(rule.Selector.Type.GetValidResultType()); parser != nil {
			result += parser.GetValue(parser.Data, rule)
		}
	}
	return
}

func (t TentacleResult) GetList(item Item) (result []string) {
	result = make([]string, 0)
	if item.Type != ItemTypeList {
		return
	}
	if len(item.Rules) < 1 {
		return
	}
	var parser *ResultParser
	for _, rule := range item.Rules {
		parser = t.GetParser(rule.Selector.Type.GetValidResultType())
		if parser == nil {
			continue
		}
		for i, val := range parser.GetValues(parser.Data, rule) {
			if len(result) <= i {
				result = append(result, val)
			} else {
				result[i] += val
			}
		}
	}
	return
}

func (t TentacleResult) GetMap(item Item) (result map[string]string) {
	result = map[string]string{}
	if item.Type != ItemTypeMap {
		return
	}
	var parser *ResultParser
	for _, rule := range item.Rules {
		parser = t.GetParser(rule.Selector.Type.GetValidResultType())
		if parser == nil {
			continue
		}
		result[rule.Key] += parser.GetValue(parser.Data, rule)
	}
	return
}

func (t TentacleResult) GetMapList(item Item) (result []map[string]string) {
	result = make([]map[string]string, 0)
	if item.Type != ItemTypeMapList {
		return
	}
	var parser *ResultParser
	for _, rule := range item.Rules {
		parser = t.GetParser(rule.Selector.Type.GetValidResultType())
		if parser == nil {
			continue
		}
		for i, val := range parser.GetValues(parser.Data, rule) {
			if len(result) <= i {
				result = append(result, newBaseMap(item))
			}
			result[i][rule.Key] += val
		}
	}
	return
}
