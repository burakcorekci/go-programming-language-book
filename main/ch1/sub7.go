package ch1

import (
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"
)

func Sub7() {
	fmt.Print(" - Chapter 1.7\n\n")
	http.HandleFunc("/", handler)
	http.HandleFunc("/count", counter)
	http.HandleFunc("/lissajous", lissajous)

	server := &http.Server{Addr: "localhost:8000", Handler: nil}
	go server.ListenAndServe()
	time.Sleep(2 * time.Second)
	server.Close()
}

var mu sync.Mutex
var count int

func handler(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	count++
	mu.Unlock()
	fmt.Fprintf(w, "%s %s %s\n", r.Method, r.URL, r.Proto)
	for k, v := range r.Header {
		fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
	}
	fmt.Fprintf(w, "Host = %q\n", r.Host)
	fmt.Fprintf(w, "RemoteAddr = %q\n", r.RemoteAddr)
	if err := r.ParseForm(); err != nil {
		log.Print(err)
	}
	for k, v := range r.Form {
		fmt.Fprintf(w, "Form[%q] = %q\n", k, v)
	}
}

func counter(w http.ResponseWriter, _ *http.Request) {
	mu.Lock()
	fmt.Fprintf(w, "Count %d\n", count)
	mu.Unlock()
}

func lissajous(w http.ResponseWriter, r *http.Request) {
	DefaultLissajous(w)
}
