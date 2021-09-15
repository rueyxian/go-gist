package main

import (
	"fmt"
	"net/http"
	"sync"
)

/*
ServeMux is an HTTP request multiplexer.
It matches the URL of each incoming request against a list of
registered patterns and calls the handler for the pattern
that most closely matches the URL.

Remark: noted that all the fields are unexported

type ServeMux struct {
	mu    sync.RWMutex
	m     map[string]muxEntry
	es    []muxEntry // slice of entries sorted from longest to shortest.
	hosts bool       // whether any patterns contain hostnames
}

type muxEntry struct {
	h       Handler
	pattern string
}
*/

type countHandler struct {
	mu    sync.Mutex
	count int
}

func (h *countHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.mu.Lock()
	defer h.mu.Unlock()
	h.count++
	fmt.Fprintf(w, "url: %s \n visit counter %d \n", r.URL.Path, h.count)
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", HelloHandlerFunc)
	mux.Handle("/url-a", &countHandler{})
	mux.Handle("/url-b", new(countHandler))

	http.ListenAndServe(":8000", mux)
}

func HelloHandlerFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, this is : %s", r.URL.Path)
}
