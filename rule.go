package deepcolor

type Item struct {
	Type  ItemType   `json:"type"`
	Rules []ItemRule `json:"rules"`
}

type ItemType string

const (
	ItemTypeSingle  ItemType = "single"
	ItemTypeList    ItemType = "list"
	ItemTypeMap     ItemType = "map"
	ItemTypeMapList ItemType = "maplist"
)

type ItemRule struct {
	Key          string            `json:"key"`
	Selector     Selector          `json:"selector"`
	Substitution map[string]string `json:"substitution"`
}

type SelectorType int

var selectorApplicableMap = map[SelectorType]ResultType{
	SelectorTypeHTMLInnerText: ResultTypeHTMl,
	SelectorTypeHTMLAttribute: ResultTypeHTMl,
	SelectorTypeTextRegExp:    ResultTypeText,
}

func (s SelectorType) GetValidResultType() ResultType {
	return selectorApplicableMap[s]
}

const (
	SelectorTypeHTMLInnerText SelectorType = 0
	SelectorTypeHTMLAttribute SelectorType = 1

	SelectorTypeTextRegExp SelectorType = 2
)

type Selector struct {
	Type  SelectorType
	Key   string
	Value string
}

func TextSelector(selector string) Selector {
	return Selector{
		Type: SelectorTypeHTMLInnerText,
		Key:  selector,
	}
}

func AttributeSelector(selector string, attribute string) Selector {
	return Selector{
		Type:  SelectorTypeHTMLInnerText,
		Key:   selector,
		Value: attribute,
	}
}

func RegExpSelector(selector string) Selector {
	return Selector{
		Type: SelectorTypeTextRegExp,
		Key:  selector,
	}
}
