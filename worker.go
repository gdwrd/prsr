package main

import (
	"fmt"
	"net/http"
	"strings"
	"sync"

	"golang.org/x/net/html"
)

type Worker struct {
	Seen    map[string]bool
	Results map[string]int
	wg      sync.WaitGroup
	Element string
}

func NewWorker(tagName string) *Worker {
	return &Worker{
		Seen:    make(map[string]bool),
		Results: make(map[string]int),
		Element: tagName,
	}
}

func (w *Worker) Start(link *Link) {
	fmt.Println("fetching: ", link.URI)
	w.Seen[link.URI] = true
	elementCount := 0
	resp, err := http.Get(link.URI)

	if err != nil {
		w.wg.Done()
		return
	}

	body := resp.Body
	defer body.Close()

	z := html.NewTokenizer(body)

	for {
		tt := z.Next()

		if len(w.Results) >= 50 {
			w.wg.Done()
			return
		}

		switch {
		case tt == html.ErrorToken:
			if len(w.Results) < 50 {
				w.Results[link.URI] = elementCount
			}

			w.wg.Done()
			return
		case tt == html.StartTagToken:
			tag := z.Token()

			if tag.Data == "a" {
				ok, url := getHref(tag)
				if !ok {
					continue
				}

				if strings.Index(url, "http") == 0 && len(w.Seen) <= 50 {
					ch <- &Link{URI: url, Level: link.Level + 1}
				}
			} else if tag.Data == w.Element {
				elementCount++
			}
		}
	}
}

func getHref(t html.Token) (ok bool, href string) {
	for _, a := range t.Attr {
		if a.Key == "href" {
			href = a.Val
			ok = true
		}
	}

	return
}
