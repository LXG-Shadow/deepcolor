package deepcolor

type SelectorType int

const (
	SelectorTypeHTMLInnerText SelectorType = 0
	SelectorTypeHTMLAttribute SelectorType = 1

	SelectorTypeTextRegExp SelectorType = 2
	SelectorTypeJsonValue  SelectorType = 3
)

type Selector struct {
	Type  SelectorType
	Path  string
	Value string
	Array bool
}

func TextSelector(selector string) *Selector {
	return &Selector{
		Type: SelectorTypeHTMLInnerText,
		Path: selector,
	}
}

func TextSliceSelector(selector string) *Selector {
	return &Selector{
		Type:  SelectorTypeHTMLInnerText,
		Path:  selector,
		Array: true,
	}
}

func AttributeSelector(selector string, attribute string) *Selector {
	return &Selector{
		Type:  SelectorTypeHTMLAttribute,
		Path:  selector,
		Value: attribute,
	}
}
func AttributeSliceSelector(selector string, attribute string) *Selector {
	return &Selector{
		Type:  SelectorTypeHTMLAttribute,
		Path:  selector,
		Value: attribute,
		Array: true,
	}
}

func RegExpSelector(selector string) *Selector {
	return &Selector{
		Type: SelectorTypeTextRegExp,
		Path: selector,
	}
}

func JsonSelector(selector string) *Selector {
	return &Selector{
		Type: SelectorTypeJsonValue,
		Path: selector,
	}
}

func JsonSliceSelector(selector string) *Selector {
	return &Selector{
		Type:  SelectorTypeJsonValue,
		Path:  selector,
		Array: true,
	}
}
