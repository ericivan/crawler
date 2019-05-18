package parser

import (
	"ericivan/crawler/engine"
	"github.com/PuerkitoBio/goquery"
	"bytes"
	"strings"
	"ericivan/crawler/model"
	"regexp"
	"strconv"
	"fmt"
)

func ParseProfile(content []byte) engine.ParseResult {

	result := engine.ParseResult{}

	profile := model.Profile{}

	reader := bytes.NewReader(content)

	document, _ := goquery.NewDocumentFromReader(reader)

	selection := document.Find(".info")

	name := selection.Find("h1[class=nickName]").Text()

	id := strings.Split(selection.Find(".id").Text(), "：")[1]

	infoString := selection.Find(".des.f-cl").Text()

	infoSlice := strings.Split(infoString, "|")

	job := infoSlice[0]

	age := extractNum(infoSlice[1])

	educ := infoSlice[2]

	marrage := infoSlice[3]

	height := extractNum(infoSlice[4])

	profile.Name = name
	profile.Age = age
	profile.Marriage = marrage
	profile.Education = educ
	profile.Height = height
	profile.Job = job
	profile.Id = id

	fmt.Printf("%+v", profile)
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
