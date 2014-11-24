package main

import (
	"fmt"
	"net/http"

	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"

	log "github.com/llimllib/loglevel"
)

func LinkReader(resp *http.Response, depth int) []Link {
	page := html.NewTokenizer(resp.Body)
	links := []Link{}

	var start *html.Token
	var text string

	for {
		_ = page.Next()
		token := page.Token()
		if token.Type == html.ErrorToken {
			break
		}

		if start != nil && token.Type == html.TextToken {
			text = fmt.Sprintf("%s%s", text, token.Data)
		}

		if token.DataAtom == atom.A {
			switch token.Type {
			case html.StartTagToken:
				if len(token.Attr) > 0 {
					start = &token
				}
			case html.EndTagToken:
				if start == nil {
					log.Warnf("Link end found without start: %s", text)
					continue
				}
				link := NewLink(*start, text, depth)
				if link.Valid() {
					links = append(links, link)
					log.Debugf("Link found: %v", link)
				}

				start = nil
				text = ""
			}
		}
	}

	log.Debug(links)

	return links
}
