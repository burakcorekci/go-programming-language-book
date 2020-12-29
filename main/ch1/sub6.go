package ch1

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"time"
)

func Sub6() {
	fmt.Print(" - Chapter 1.6\n\n")
	fetchAll([]string{
		"https://raw.githubusercontent.com/burakcorekci/go-programming-language-book/master/Dockerfile",
		"https://raw.githubusercontent.com/burakcorekci/go-programming-language-book/master/build.gradle",
	})
}

func fetchAll(urls []string) {
	start := time.Now()
	ch := make(chan string)
	for _, url := range urls {
		go fetch(url, ch) // start a goroutine
	}
	for range urls {
		fmt.Println(<-ch) // receive from channel ch
	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetch(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err) // send to channel ch
		return
	}

	nbytes, err := io.Copy(ioutil.Discard, resp.Body)
	resp.Body.Close() // don't leak resources
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs %7d %s", secs, nbytes, url)
}
