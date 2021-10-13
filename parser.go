package deepcolor

import (
	"github.com/PuerkitoBio/goquery"
	"regexp"
)

func subText(content string, rule ItemRule) string {
	if rule.Substitution == nil {
		return content
	}
	for key, val := range rule.Substitution {
		content = regexp.MustCompile(key).ReplaceAllString(content, val)
	}
	return content
}

func subHTML(selection *goquery.Selection, rule ItemRule) *goquery.Selection {
	if rule.Substitution == nil {
		return selection
	}
	htmltext, _ := selection.Html()
	for key, val := range rule.Substitution {
		htmltext = regexp.MustCompile(key).ReplaceAllString(htmltext, val)
	}
	return selection.SetHtml(htmltext)
}

func newBaseMap(collection Item) map[string]string {
	baseMap := map[string]string{}
	for _, rule := range collection.Rules {
		baseMap[rule.Key] = ""
	}
	return baseMap
}

func getTextValue(i interface{}, rule ItemRule) (value string) {
	t := i.(TentacleTextResult)
	if rule.Selector == "" {
		return ""
	}
	return subText(regexp.MustCompile(rule.Selector).FindString(t.Data.(string)), rule)
}

func getTextValues(i interface{}, rule ItemRule) (values []string) {
	t := i.(TentacleTextResult)
	values = make([]string, 0)
	if rule.Selector == "" {
		return
	}
	for _, val := range regexp.MustCompile(rule.Selector).FindAllString(t.Data.(string), -1) {
		values = append(values, subText(val, rule))
	}
	return
}

func getHTMLValue(i interface{}, rule ItemRule) string {
	t := i.(TentacleTextResult).Data
	if rule.Selector == "" {
		return ""
	}
	switch rule.Target.Type {
	case SelectorTargetTypeHTMLInnerText:
		return subHTML(t.(*goquery.Document).Find(rule.Selector), rule).Text()
	case SelectorTargetTypeHTMLAttribute:
		attr, _ := subHTML(t.(*goquery.Document).Find(rule.Selector), rule).Attr(rule.Target.Value)
		attr = subText(attr, rule)
		return attr
	default:
		text, _ := t.(*goquery.Document).Html()
		return subText(regexp.MustCompile(rule.Selector).FindString(text), rule)
	}
}

func getHTMLValues(i interface{}, rule ItemRule) (values []string) {
	t := i.(TentacleTextResult).Data
	values = make([]string, 0)
	if rule.Selector == "" {
		return
	}
	switch rule.Target.Type {
	case SelectorTargetTypeHTMLInnerText:
		t.(*goquery.Document).Find(rule.Selector).Each(func(i int, selection *goquery.Selection) {
			values = append(values, subHTML(selection, rule).Text())
		})
	case SelectorTargetTypeHTMLAttribute:
		t.(*goquery.Document).Find(rule.Selector).Each(func(i int, selection *goquery.Selection) {
			attr, _ := subHTML(selection, rule).Attr(rule.Target.Value)
			attr = subText(attr, rule)
			values = append(values, attr)
		})
	default:
		text, _ := t.(*goquery.Document).Html()
		for _, val := range regexp.MustCompile(rule.Selector).FindAllString(text, 0) {
			values = append(values, subText(val, rule))
		}
	}
	return
}

// todo json same key append
//
//func ParseJsonSingle(doc *gjson.Result, item Item) (result string) {
//	result = ""
//	if item.Type != ItemTypeSingle {
//		return
//	}
//	if len(item.Rules) != 1 {
//		return
//	}
//	result = doc.Get(item.Rules[0].Selector).String()
//	return
//}
//
//func ParseJsonList(doc *gjson.Result, collection Item) (result []string) {
//	result = make([]string, 0)
//	if collection.Type != ItemTypeList {
//		return
//	}
//	if len(collection.Rules) != 1 {
//		return
//	}
//	rule := collection.Rules[0]
//	doc.Get(rule.Selector).ForEach(func(key, value gjson.Result) bool {
//		result = append(result, value.String())
//		return true
//	})
//	return
//}
//
//func ParseJsonMap(doc *gjson.Result, collection Item) (result map[string]string) {
//	result = map[string]string{}
//	if collection.Type != ItemTypeMap {
//		return
//	}
//	for _, rule := range collection.Rules {
//		result[rule.Key] = doc.Get(rule.Selector).String()
//	}
//	return
//}
//
//func ParseJsonMapList(doc *gjson.Result, collection Item) (result []map[string]string) {
//	result = make([]map[string]string, 0)
//	if collection.Type != ItemTypeMapList {
//		return
//	}
//	for _, rule := range collection.Rules {
//		index := 0
//		doc.Get(rule.Selector).ForEach(func(key, value gjson.Result) bool {
//			if len(result) <= index {
//				result = append(result, newBaseMap(collection))
//			}
//			result[index][rule.Selector] = value.String()
//			return true
//		})
//	}
//	return
//}
