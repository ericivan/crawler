package engine

import (
	"ericivan/crawler/fetcher"
	"log"
	"errors"
)

type SimpleEngine struct {

}

func (e SimpleEngine)Run(seeds ...Request) {
	var requests []Request

	for _, r := range seeds {
		requests = append(requests, r)
	}

	for len(requests) > 0 {

		r := requests[0]
		requests = requests[1:]

		parseResult, err := Worker(r)

		if err != nil {
			log.Printf("error is %s", err.Error())
			continue
		}

		requests = append(requests, parseResult.Request...)

		for _, item := range parseResult.Items {
			log.Printf("Got Item %+v \n ", item)
		}

	}
}
func Worker(r Request) (ParseResult, error) {
	if r.Url == "" {
		return ParseResult{}, errors.New("null Url")
	}

	log.Printf("Fetching %s", r.Url)

	body, err := fetcher.Fetch(r.Url)

	if err != nil {
		log.Printf("Fetcher err: %s,%s", r.Url, err.Error())
		return ParseResult{}, errors.New("null Url")
	}

	return r.ParserFunc(body), nil
}
