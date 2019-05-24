package parser

import (
	"bytes"
	"github.com/PuerkitoBio/goquery"
	"fmt"
	"ericivan/crawler/engine"
)

func ParseMoveList(content []byte) engine.ParseResult {

	reader := bytes.NewReader(content)

	result := engine.ParseResult{}

	dom, err := goquery.NewDocumentFromReader(reader)

	if err != nil {
		fmt.Println(err.Error)

		return result
	}

	dom.Find(".post.clearfix").Each(func(i int, s *goquery.Selection) {

		url, _ := s.Find(".postli-1 a").Attr("href")

		title := s.Find(".postli-1 a").Text()

		result.Request = append(result.Request, engine.Request{
			Url:        url,
			ParserFunc: ParserMovie,
		})
		result.Items = append(result.Items, title)
	})

	nextUrl, ok := dom.Find(".nextpostslink").Attr("href")

	if (ok) {
		result.Request = append(result.Request, engine.Request{
			Url:        nextUrl,
			ParserFunc: ParseMoveList,
		})
	}

	return result

}
