package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

/*

func Handle(pattern string, handler Handler)

type Handler interface {
    ServeHTTP(ResponseWriter, *Request)
}

*/

type countHandler struct {
	mu sync.Mutex // guards n
	n  int
}

func (h *countHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.mu.Lock()
	defer h.mu.Unlock()
	h.n++
	fmt.Fprintf(w, "count is %d\n", h.n)
}

func main() {
	// http.Handle("/count", new(countHandler))
	// log.Fatal(http.ListenAndServe(":8080", nil))

	http.Handle("/count", countHandler{})
	log.Fatal(http.ListenAndServe(":8080", nil))

}
