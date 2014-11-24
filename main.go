package main

import (
	"fmt"
	"net/http"
	"os"

	log "github.com/llimllib/loglevel"
)

var MaxDepth = 2

func main() {
	log.SetPriorityString("info")
	log.SetPrefix("crawler")

	log.Debug(os.Args)

	if len(os.Args) < 2 {
		log.Fatalln("Missing URL arg")
	}

	recurDownloader(os.Args[1], 0)
}

func recurDownloader(url string, depth int) {
	page, err := downloader(url)
	if err != nil {
		log.Error(err)
		return
	}

	links := LinkReader(page, depth)

	for _, link := range links {
		fmt.Println(link)
		if depth+1 < MaxDepth {
			recurDownloader(link.url, depth+1)
		}
	}
}

func downloader(url string) (resp *http.Response, err error) {
	log.Debugf("Downloading %s", url)

	resp, err = http.Get(url)
	if err != nil {
		log.Debugf("Error: %s", err)
		return
	}

	if resp.StatusCode > 299 {
		err = HttpError{fmt.Sprintf("Error (%d): %s", resp.StatusCode, url)}
		log.Debug(err)
		return
	}

	return
}
