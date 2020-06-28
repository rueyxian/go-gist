package main

import (
	"fmt"
	"net/http"
)

/*
ServeMux is an http request multiplexer, which responsible for matching the url in the request to an appropriate handler and executing it.

Multiplexer or mux also known as data selector, which is a mechanism of that select serveral input and forward it into a single output.

*/

/*
func (mux *ServeMux) Handle(pattern string, handler Handler)

type Handler interface {
    ServeHTTP(ResponseWriter, *Request)
}
*/

func main() {

	mux := http.NewServeMux()

	fmt.Printf("%T \n", mux)

	// mux.Handle("/", home{})

}
