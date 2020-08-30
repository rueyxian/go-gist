package main

import (
	"fmt"
	"log"
	"net/http"
)

/*
type Handler interface{
	serveHTTP(responseWriter, *Request)
}
*/

type simpleHandler struct{}

func (h *simpleHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "this is simpleHandler : %s", r.URL.Path)
}

func main() {

	log.Fatal(http.ListenAndServe(":8000", &simpleHandler{}))

}
