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

func SelectorText(selector string) *Selector {
	return &Selector{
		Type: SelectorTypeHTMLInnerText,
		Path: selector,
	}
}

func SelectorTextSlice(selector string) *Selector {
	return &Selector{
		Type:  SelectorTypeHTMLInnerText,
		Path:  selector,
		Array: true,
	}
}

func SelectorAttribute(selector string, attribute string) *Selector {
	return &Selector{
		Type:  SelectorTypeHTMLAttribute,
		Path:  selector,
		Value: attribute,
	}
}
func SelectorAttributeSlice(selector string, attribute string) *Selector {
	return &Selector{
		Type:  SelectorTypeHTMLAttribute,
		Path:  selector,
		Value: attribute,
		Array: true,
	}
}

func SelectorRegExp(selector string) *Selector {
	return &Selector{
		Type: SelectorTypeTextRegExp,
		Path: selector,
	}
}

func SelectorJson(selector string) *Selector {
	return &Selector{
		Type: SelectorTypeJsonValue,
		Path: selector,
	}
}

func SelectorJsonSlice(selector string) *Selector {
	return &Selector{
		Type:  SelectorTypeJsonValue,
		Path:  selector,
		Array: true,
	}
}
