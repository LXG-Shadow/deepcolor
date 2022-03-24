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
	if rule.Selector.Key == "" {
		return ""
	}
	return subText(regexp.MustCompile(rule.Selector.Key).FindString(i.(string)), rule)
}

func getTextValues(i interface{}, rule ItemRule) (values []string) {
	values = make([]string, 0)
	if rule.Selector.Key == "" {
		return
	}
	for _, val := range regexp.MustCompile(rule.Selector.Key).FindAllString(i.(string), -1) {
		values = append(values, subText(val, rule))
	}
	return
}

func getHTMLValue(i interface{}, rule ItemRule) string {
	t := i.(*goquery.Document)
	if rule.Selector.Key == "" {
		return ""
	}
	switch rule.Selector.Type {
	case SelectorTypeHTMLInnerText:
		return subHTML(t.Find(rule.Selector.Key), rule).Text()
	case SelectorTypeHTMLAttribute:
		attr, _ := subHTML(t.Find(rule.Selector.Key), rule).Attr(rule.Selector.Value)
		attr = subText(attr, rule)
		return attr
	default:
		text, _ := t.Html()
		return subText(regexp.MustCompile(rule.Selector.Key).FindString(text), rule)
	}
}

func getHTMLValues(i interface{}, rule ItemRule) (values []string) {
	t := i.(*goquery.Document)
	values = make([]string, 0)
	if rule.Selector.Key == "" {
		return
	}
	switch rule.Selector.Type {
	case SelectorTypeHTMLInnerText:
		t.Find(rule.Selector.Key).Each(func(i int, selection *goquery.Selection) {
			values = append(values, subHTML(selection, rule).Text())
		})
	case SelectorTypeHTMLAttribute:
		t.Find(rule.Selector.Key).Each(func(i int, selection *goquery.Selection) {
			attr, _ := subHTML(selection, rule).Attr(rule.Selector.Value)
			attr = subText(attr, rule)
			values = append(values, attr)
		})
	default:
		text, _ := t.Html()
		for _, val := range regexp.MustCompile(rule.Selector.Key).FindAllString(text, 0) {
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
