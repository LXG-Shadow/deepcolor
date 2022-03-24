package deepcolor

import "github.com/tidwall/gjson"

func getJSONValue(i interface{}, rule ItemRule) string {
	g := i.(gjson.Result)
	if rule.Selector.Key == "" {
		return ""
	}
	return subText(g.Get(rule.Selector.Key).String(), rule)
}

func getJSONValues(i interface{}, rule ItemRule) (result []string) {
	g := i.(gjson.Result)
	if rule.Selector.Key == "" {
		return
	}
	for _, name := range g.Get(rule.Selector.Key).Array() {
		result = append(result, subText(name.String(), rule))
	}
	return
}
