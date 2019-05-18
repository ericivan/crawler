package engine

import (
	"ericivan/crawler/fetcher"
	"log"
)

func Run(seeds ...Request) {
	var requests []Request

	for _, r := range seeds {
		requests = append(requests, r)
	}

	for len(requests) > 0 {

		r := requests[0]
		requests = requests[1:]


		if r.Url == "" {
			continue
		}

		log.Printf("Fetching %s", r.Url)

		body, err := fetcher.Fetch(r.Url)

		if err != nil {
			log.Printf("Fetcher err: %s,%s", r.Url, err.Error())
			continue
		}

		parseResult := r.ParserFunc(body)

		requests = append(requests, parseResult.Request...)

		for _, item := range parseResult.Items {
			log.Printf("Got Item %+v \n ", item)
		}

	}
}
