package engine

import (
	"log"
	"soy/fetcher"
)

func Run(seed ...Request) {
	var requests []Request

	for _, request := range seed {
		requests = append(requests, request)
	}

	for len(requests) > 0 {
		r := requests[0]
		requests = requests[1:]

		result, err := worker(r)
		if err != nil {
			continue
		}
		requests = append(requests, result.Requests...)

		for _, item := range result.Items {
			log.Printf("Got item, %#v", item)
		}
	}

}

func worker(r Request) (ParserResult, error) {
	body, err := fetcher.Fetcher(r.Url)
	log.Printf("%s",body)
	if err != nil {
		log.Printf("Fetcher: error fetching url %s: %v", r.Url, err)
		return ParserResult{}, nil
	}

	return r.ParserFunc(body), nil
}
