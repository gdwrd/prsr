package main

import (
	"testing"

	"golang.org/x/net/html"
)

func TestNewWorker(t *testing.T) {
	w := NewWorker(&Config{
		TagName: "test",
	})

	if w.Conf.TagName != "test" {
		t.Error("TagName should be equal")
	}
}

func TestParseLinkTag(t *testing.T) {
	tag := html.Token{
		Attr: make([]html.Attribute, 1),
		Data: "a",
	}

	tag.Attr = append(tag.Attr, html.Attribute{Key: "href", Val: "http://sheremet.pw"})

	_, url := parseLinkTag(tag)

	if url != "http://sheremet.pw" {
		t.Error("URL should be equal")
	}
}
