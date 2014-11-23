package main

import (
	"strings"
	"testing"
)

func TestLinkStringPrintsDepth(t *testing.T) {
	prefixes := []string{"", "\t", "\t\t"}

	for i := 0; i < 3; i++ {
		link := Link{depth: i, text: "Key", url: "value"}

		str := link.String()

		if !strings.HasPrefix(str, prefixes[i]) {
			t.Errorf("Wrong Prefix for depth %d: '%s'", link.depth, str)
		}
	}
}
