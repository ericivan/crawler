package parser

import (
	"ericivan/crawler/engine"
	"github.com/PuerkitoBio/goquery"
	"bytes"
	"fmt"
	"strings"
)

func ParseProfile(content []byte) engine.ParseResult {

	result := engine.ParseResult{}

	reader := bytes.NewReader(content)

	document, _ := goquery.NewDocumentFromReader(reader)

	selection := document.Find("").Last()

	fmt.Println(strings.Repeat("-", 100))

	fmt.Println(selection.Html())
	name := selection.Find(".name .nickName").Text()

	id := selection.Find(".id").Text()

	infoString := selection.Find(".des.f-cl").Text()

	fmt.Printf("infoString is %s \n", infoString)
	fmt.Printf("ID is %s \n", id)
	fmt.Println(selection.Find("a").Html())

	fmt.Printf("name is %s \n", name)

	fmt.Println(strings.Repeat("-", 100))
	return result

}
