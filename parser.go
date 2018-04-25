package main

import (
	"flag"
	"fmt"
)

var conf *Config

// init function
//
// Parsing arguments using flag
func init() {
	uri := flag.String("uri", "https://sheremet.pw/", "URI to parse")
	tag := flag.String("tag", "input", "Tag name you want to find")
	deep := flag.Int("d", 3, "Parser page level deep")
	maxSize := flag.Int("l", 50, "Max size of parsed links")

	flag.Parse()

	conf = &Config{
		BaseURI: *uri,
		TagName: *tag,
		MaxDeep: *deep,
		MaxLink: *maxSize,
	}
}

// main function
func main() {
	worker := NewWorker(conf)

	parse(worker)
}

// parse function
//
// Start to parsing baseLink
//
// Params:
// - w {*Worker}
func parse(w *Worker) {
	baseLink := &Link{URI: conf.BaseURI}

	w.wg.Add(1)
	go w.Start(baseLink)

	go func() {
		for item := range w.Channel {
			if item.Level < w.Conf.MaxDeep && !w.Seen.Exist(item.URI) {
				w.wg.Add(1)
				go w.Start(item)
			}
		}
	}()

	w.wg.Wait()

	fmt.Printf("\n")

	for k, v := range w.Results.data {
		fmt.Println("URI: ", k, "has", v, w.Conf.TagName)
	}

	fmt.Println("\nLINKS PARSED: ", len(w.Results.data))
}
