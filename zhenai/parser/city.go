package parser

import (
	"bytes"
	"github.com/PuerkitoBio/goquery"
	"ericivan/crawler/engine"
)

func ParseCity(c []byte) engine.ParseResult {

	result := engine.ParseResult{}

	reader := bytes.NewReader(c)

	dom, err := goquery.NewDocumentFromReader(reader)

	if err != nil {

	}

	dom.Find(".g-list .list-item").Each(func(i int, s *goquery.Selection) {
		url, _ := s.Find("div[class=photo] a").Attr("href")

		result.Request = append(result.Request, engine.Request{
			Url:        url,
			ParserFunc: ParseProfile,
		})

		result.Items = append(result.Items, url)

	})

	return result
}
