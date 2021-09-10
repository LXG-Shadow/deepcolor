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
	Selector     string            `json:"selector"`
	Target       SelectorTarget    `json:"target"`
	Substitution map[string]string `json:"substitution"`
}

type SelectorTargetType string

const (
	SelectorTargetTypeHTMLInnerText SelectorTargetType = "html_innertext"
	SelectorTargetTypeHTMLAttribute SelectorTargetType = "html_attribute"

	SelectorTargetTypeTextRegExp SelectorTargetType = "text_regexp"

	//SelectorTargetTypeJson SelectorTargetType = "json"
)

type SelectorTarget struct {
	Type  SelectorTargetType
	Value string
}

func TextTarget() SelectorTarget {
	return SelectorTarget{SelectorTargetTypeHTMLInnerText, ""}
}

func AttributeTarget(attribute string) SelectorTarget {
	return SelectorTarget{SelectorTargetTypeHTMLAttribute, attribute}
}

func RegExpTarget() SelectorTarget {
	return SelectorTarget{SelectorTargetTypeTextRegExp, ""}
}
//
//func JsonTarget() SelectorTarget {
//	return SelectorTarget{SelectorTargetTypeJson, ""}
//}
