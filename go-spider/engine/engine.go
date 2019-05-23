package engine

import (
	"fmt"
	"go-spider/fetcher"
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
		log.Printf("Fetching %s\n", r.Url)
		body, err := fetcher.Fetch(r.Url)
		if err != nil {
			log.Printf("Fetcher: error fetching url %s:%v", r.Url, err)
			continue
		}
		ParseResult := r.ParserFunc(body)
		requests = append(requests, ParseResult.Requests...)
		for _, item := range ParseResult.Items {
			fmt.Printf("Got item %v\n", item)
		}
	}
}
