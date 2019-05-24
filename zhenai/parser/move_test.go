package parser

import (
	"testing"
	"ericivan/crawler/fetcher"
	"fmt"
)

func TestParseTime(t *testing.T) {
	ParseTime("hahah")

}

func TestParseMoveCategory(t *testing.T) {

	url := "https://www.xl720.com/category/dongzuopian"

	content, err := fetcher.Fetch(url)

	if err != nil {
		fmt.Println(err.Error())
	}

	ParseMoveList(content)
}

func TestParserMove(t *testing.T) {
	url:= "https://www.xl720.com/thunder/5842.html"

	content, err := fetcher.Fetch(url)

	if err != nil {
		fmt.Println(err.Error())
	}

	ParserMovie(content)

}
