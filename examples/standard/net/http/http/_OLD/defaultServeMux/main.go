package main

import (
	"fmt"
	"net/http"
	"sync"
)

/*
// DefaultServeMux is the default ServeMux used by Serve.
var DefaultServeMux = &defaultServeMux

var defaultServeMux ServeMux

// Handle registers the handler for the given pattern
// in the DefaultServeMux.
// The documentation for ServeMux explains how patterns are matched.
func Handle(pattern string, handler Handler) { DefaultServeMux.Handle(pattern, handler) }

// HandleFunc registers the handler function for the given pattern
// in the DefaultServeMux.
// The documentation for ServeMux explains how patterns are matched.
func HandleFunc(pattern string, handler func(ResponseWriter, *Request)) {
	DefaultServeMux.HandleFunc(pattern, handler)
}

func (sh serverHandler) ServeHTTP(rw ResponseWriter, req *Request) {
	handler := sh.srv.Handler
	if handler == nil {
		handler = DefaultServeMux
	}
	if req.RequestURI == "*" && req.Method == "OPTIONS" {
		handler = globalOptionsHandler{}
	}
	handler.ServeHTTP(rw, req)
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

	// http package has a built-in ServeMux, which is called DefaultServeMux
	// and also provides HandleFunc() and Handle() as well

	http.HandleFunc("/", HelloHandlerFunc)
	http.Handle("/url-a", &countHandler{})
	http.Handle("/url-b", new(countHandler))

	// to user http.DefaultServeMux, pass nil as http.Handler argument
	http.ListenAndServe(":8000", nil)
}

func HelloHandlerFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, this is : %s", r.URL.Path)
}
