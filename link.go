package main

import (
	"fmt"
	"strings"
)

type Link struct {
	url   string
	text  string
	depth int
}

func (self Link) String() string {
	spacer := strings.Repeat("\t", self.depth)
	return fmt.Sprintf("%s%s - %s\n", spacer, self.text, self.url)
}
