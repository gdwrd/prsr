package main

import (
	"testing"
)

func TestSeenMapAdd(t *testing.T) {
	m := &SeenMap{data: make(map[string]bool)}
	m.Add("http://test.com/")

	if !m.data["http://test.com/"] {
		t.Error("Link should be added")
	}
}

func TestSeenMapExist(t *testing.T) {
	m := &SeenMap{data: make(map[string]bool)}
	m.Add("http://test.com/")

	if !m.Exist("http://test.com/") {
		t.Error("Link should exist")
	}
}

func TestResultsMapAdd(t *testing.T) {
	m := &ResultsMap{data: make(map[string]int)}
	m.Add("http://test.com/", 10)

	if m.data["http://test.com/"] != 10 {
		t.Error("Value should be as expected")
	}
}
