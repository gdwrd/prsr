package main

import (
	"flag"
	"fmt"
	"io"
)

type Link struct {
	URI   string
	Data  io.ReadCloser
	Level int
}

var ch = make(chan *Link, 10)

func main() {
	uri := flag.String("uri", "https://sheremet.pw/", "URI to parse")
	tag := flag.String("tag", "input", "Tag name you want to find")
	flag.Parse()

	w := NewWorker(*tag)
	baseLink := &Link{URI: *uri}

	w.wg.Add(1)
	go w.Start(baseLink)

	go func() {
		for item := range ch {
			if item.Level < 3 && !w.Seen[item.URI] {
				w.wg.Add(1)
				go w.Start(item)
			}

			if len(w.Results) >= 50 || len(w.Seen) >= 50 {
				break
			}
		}
	}()

	w.wg.Wait()
	fmt.Println("LINKS PARSED: ", len(w.Results))

	for k, v := range w.Results {
		fmt.Println("URI: ", k, "has", v, w.Element)
	}
}
