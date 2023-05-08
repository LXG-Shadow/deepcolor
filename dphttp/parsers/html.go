package parsers

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/aynakeya/deepcolor/dphttp"
	"strings"
)

func HTMLGoqueryParser(resp *dphttp.Response) (*goquery.Document, error) {
	return goquery.NewDocumentFromReader(strings.NewReader(resp.String()))
}
