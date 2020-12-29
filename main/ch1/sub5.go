package ch1

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

func Sub5() {
	fmt.Print(" - Chapter 1.5\n\n")
	printContentsFromUrls([]string{"raw.githubusercontent.com/burakcorekci/go-programming-language-book/master/Dockerfile"})
}

func printContentsFromUrls(urls []string) {
	for _, url := range urls {
		b, err := getContentFromUrl(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
		}
		fmt.Printf("%s\n\n", b)
	}
}

func getContentFromUrl(url string) ([]byte, error) {
	url = normalizeUrl(url)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	b, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	fmt.Printf(" -- %s, Response Code: %v\n", url, resp.StatusCode)
	if err != nil {
		return nil, err
	}
	return b, nil
}

func normalizeUrl(url string) string {
	if !(strings.HasPrefix(url, "http://") || strings.HasPrefix(url, "https://")) {
		url = "http://" + url
	}
	return url
}
