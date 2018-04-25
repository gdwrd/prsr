package main

import (
	"fmt"
	"net/http"
	"strings"
	"sync"

	"golang.org/x/net/html"
)

// Worker struct
//
// Main worker struct
// Used for parse links and tags
type Worker struct {
	Seen    *SeenMap
	Results *ResultsMap
	wg      sync.WaitGroup
	Conf    *Config
	Channel chan *Link
}

// NewWorker function
//
// Create new worker with config
//
// Params:
// - config {*Config}
//
// Result:
// - {*Worker}
func NewWorker(config *Config) *Worker {
	return &Worker{
		Seen:    &SeenMap{data: make(map[string]bool)},
		Results: &ResultsMap{data: make(map[string]int)},
		Conf:    config,
		Channel: make(chan *Link),
	}
}

// Start function
//
// Used for start worker
//
// Params:
// - link {*Link}
func (w *Worker) Start(link *Link) {
	fmt.Print(".")
	w.Seen.Add(link.URI)
	resp, err := http.Get(link.URI)

	if err != nil {
		fmt.Println("Connection error with ", link.URI)
		w.wg.Done()
		return
	}

	link.Data = resp.Body
	defer resp.Body.Close()

	w.ParseBody(link)
}

// ParseBody function
//
// Parsing body, add new links to w.Channel
// Counting w.Conf.TagName
//
// Params:
// - link {*Link}
func (w *Worker) ParseBody(link *Link) {
	data := html.NewTokenizer(link.Data)
	elementCount := 0

	for {
		value := data.Next()

		if len(w.Results.data) >= w.Conf.MaxLink {
			w.wg.Done()
			break
		}

		switch {
		case value == html.ErrorToken:
			if len(w.Results.data) < w.Conf.MaxLink {
				w.Results.Add(link.URI, elementCount)
			}

			w.wg.Done()
			return
		case value == html.StartTagToken:
			tag := data.Token()

			if tag.Data == "a" {
				ok, url := parseLinkTag(tag)
				if !ok {
					continue
				}

				if len(w.Seen.data) <= w.Conf.MaxLink {
					w.Channel <- &Link{URI: url, Level: link.Level + 1}
				}
			} else if tag.Data == w.Conf.TagName {
				elementCount++
			}
		}
	}
}

// parseLinkTag function
//
// Check if link is valid and has valid url
//
// Params:
// - t {html.Token}
//
// Result:
// - ok 	{bool}
// - href {string}
func parseLinkTag(t html.Token) (ok bool, href string) {
	for _, a := range t.Attr {
		if a.Key == "href" && strings.Index(a.Val, "http") == 0 {
			href = a.Val
			ok = true
		}
	}

	return
}
