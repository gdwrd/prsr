package main

import (
	"testing"
)

func TestParse(t *testing.T) {
	worker := getDefaultWorker()
	parse(worker)

	if len(worker.Results.data) == 0 {
		t.Error("Results data shouldn't be empty")
	}
}

func getDefaultWorker() *Worker {
	return NewWorker(&Config{
		BaseURI: "https://sheremet.pw",
		MaxDeep: 2,
		MaxLink: 2,
		TagName: "input",
	})
}
