package deepcolor

import (
	"github.com/PuerkitoBio/goquery"
	"strings"
)

type ParserHTML struct {
	doc *goquery.Document
}

func NewHTMLParser() ResponseParser {
	return &ParserHTML{}
}

func (p *ParserHTML) Initialize(resp *Response) (err error) {
	//p.doc, err = goquery.NewDocumentFromReader(strings.NewReader(dputil.DecodeString(resp.String(), resp.Request.Charset)))
	p.doc, err = goquery.NewDocumentFromReader(strings.NewReader(resp.String()))
	return err
}

func (p *ParserHTML) Get(selector *Selector) interface{} {
	if selector.Array {
		return p.GetValues(selector)
	}
	return p.GetValue(selector)
}

func (p *ParserHTML) GetValue(selector *Selector) interface{} {
	if selector.Path == "" {
		return ""
	}
	switch selector.Type {
	case SelectorTypeHTMLInnerText:
		return p.doc.Find(selector.Path).Text()
	case SelectorTypeHTMLAttribute:
		attr, _ := p.doc.Find(selector.Path).Attr(selector.Value)
		return attr
	}
	return ""
}

func (p *ParserHTML) GetValues(selector *Selector) []interface{} {
	values := make([]interface{}, 0)
	if selector.Path == "" {
		return values
	}
	switch selector.Type {
	case SelectorTypeHTMLInnerText:
		p.doc.Find(selector.Path).Each(func(i int, selection *goquery.Selection) {
			values = append(values, selection.Text())
		})
	case SelectorTypeHTMLAttribute:
		p.doc.Find(selector.Path).Each(func(i int, selection *goquery.Selection) {
			attr, _ := selection.Attr(selector.Value)
			values = append(values, attr)
		})
	}
	return values
}
