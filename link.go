package main

import (
	"fmt"
	"strings"

	"golang.org/x/net/html"
)

type Link struct {
	url   string
	text  string
	depth int
}

func NewLink(tag html.Token, text string, depth int) Link {
	link := Link{text: strings.TrimSpace(text), depth: depth}

	for i := range tag.Attr {
		if tag.Attr[i].Key == "href" {
			link.url = strings.TrimSpace(tag.Attr[i].Val)
		}
	}

	return link
}

func (self Link) String() string {
	spacer := strings.Repeat("\t", self.depth)
	return fmt.Sprintf("%s%s (%d) - %s", spacer, self.text, self.depth, self.url)
}

func (self Link) Valid() bool {
	if self.depth >= MaxDepth {
		return false
	}

	if len(self.text) == 0 {
		return false
	}

	if len(self.url) == 0 || strings.Contains(strings.ToLower(self.url), "javascript") {
		return false
	}

	return true
}
