package parser

import (
	"ericivan/crawler/engine"
	"bytes"
	"github.com/PuerkitoBio/goquery"
	"strings"
	"ericivan/crawler/model"
	"regexp"
	"strconv"
	"time"
	"fmt"
)

func ParserMovie(content []byte) engine.ParseResult {
	result := engine.ParseResult{}

	reader := bytes.NewReader(content)

	dom, _ := goquery.NewDocumentFromReader(reader)

	Movie := model.MoveDetail{}

	title := dom.Find(".current").Text()
	Movie.Title = title

	dom.Find(".meta-author a").Each(func(i int, s *goquery.Selection) {
		Movie.Type = append(Movie.Type, strings.Trim(s.Text(), ""))
	})

	dom.Find("#info [href*='https://www.xl720.com/thunder/writers']").Each(func(i int, s *goquery.Selection) {
		Movie.Director = append(Movie.Director, s.Text())
	})

	Movie.Publish = ParseTime(dom.Find(".meta-date").Text())

	dom.Find("#info [href*='https://www.xl720.com/thunder/stars']").Each(func(i int, s *goquery.Selection) {
		Movie.Actors = append(Movie.Actors, s.Text())
	})

	Movie.Description = dom.Find("#link-report").Text()

	fmt.Printf("%+v", Movie)
	return result
}
func ParseTime(str string) time.Time {

	reg := regexp.MustCompile("([0-9]+).*([0-9]+).*([0-9]+)")

	result := reg.FindAllSubmatch([]byte(str), -1)

	fmt.Println("parser time result is ")

	sub := result[0]

	year, _ := strconv.Atoi(string(sub[1]))
	month, _ := strconv.Atoi(string(sub[2]))
	day, _ := strconv.Atoi(string(sub[3]))

	the_time := time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.Local).Unix()

	return time.Unix(the_time, 0)
}
