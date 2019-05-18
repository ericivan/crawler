package parser

import (
	"ericivan/crawler/engine"
	"github.com/PuerkitoBio/goquery"
	"bytes"
	"ericivan/crawler/model"
	"regexp"
	"strconv"
	"strings"
)

func ParseProfile(content []byte) engine.ParseResult {

	result := engine.ParseResult{}

	profile := model.Profile{}

	reader := bytes.NewReader(content)

	document, _ := goquery.NewDocumentFromReader(reader)

	document.Find(".info").Each(func(i int, selection *goquery.Selection) {
		if (i == 4) {

			name := selection.Find(".nickName").Text()
			profile.Name = name

			infoString := selection.Find(".des.f-cl").Text()

			infoSlice := strings.Split(infoString, "|")

			if len(infoSlice) == 6 {
				job := infoSlice[0]
				age := extractNum(infoSlice[1])
				educ := infoSlice[2]
				marrage := infoSlice[3]
				height := extractNum(infoSlice[4])

				profile.Age = age
				profile.Marriage = marrage
				profile.Education = educ
				profile.Height = height
				profile.Job = job
			}
		}

		IdParse := strings.Split(selection.Find(".id").Text(), "：")

		if len(IdParse) == 2 {
			id := IdParse[1]
			profile.Id = id
		}
	})


	result.Items = append(result.Items, profile)

	result.Request = append(result.Request, engine.Request{
		Url:        "",
		ParserFunc: engine.NilParseFunc,
	})
	return result
}

func extractNum(c string) int {
	num, err := strconv.Atoi(
		string(regexp.MustCompile("[0-9]+").
			Find([]byte(c))))

	if err != nil {
		return 0
	}

	return num
}
