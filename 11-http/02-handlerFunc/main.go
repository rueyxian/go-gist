package main

import (
	"fmt"
	"log"
	"net/http"
)

/*
The HandlerFunc type is an adapter to allow the use of ordinary functions as HTTP handlers.
If f is a function with the appropriate signature, HandlerFunc(f) is a Handler that calls f.

type HandlerFunc func(ResponseWriter, *Request)

func (f HandlerFunc) ServeHTTP (w ResponseWriter, r *Request){
	f(w,r)
}
*/

func simpleHandlerFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "this is simpleHandlerFunc V2 : %s", r.URL.Path)
}

func newSimpleHandler() http.Handler {
	// the signature implies that this function return http.Handler interface
	// it need to be converted into http.HandlerFunc type as http.HandlerFunc implements http.handler
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "this is simpleHandlerFunc V3 : %s", r.URL.Path)
	})
}

func newSimpleHandlerFunc() http.HandlerFunc {
	// conversion can be ommited because it returns a function which
	// it's signature matches http.HandlerFunc
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "this is simpleHandlerFunc V4 : %s", r.URL.Path)
	}
}

func main() {

	h1 := func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "this is simpleHandlerFunc v1 : %s", r.URL.Path)
	}
	h2 := simpleHandlerFunc
	h3 := newSimpleHandler()
	h4 := newSimpleHandlerFunc()

	fmt.printf("h1: %t \n", h1) // func(http.ResponseWriter, *http.Request)
	fmt.Printf("h2: %T \n", h2) // func(http.ResponseWriter, *http.Request)
	fmt.Printf("h3: %T \n", h3) // http.HandlerFunc
	fmt.Printf("h4: %T \n", h4) // http.HandlerFunc

	// for h1 and h2, even though it's signature match http.HandlerFunc,
	// conversion is required in order to be passed as an argument
	log.Fatal(http.ListenAndServe(":8000", http.HandlerFunc(h1)))
	// log.Fatal(http.ListenAndServe(":8000", http.HandlerFunc(h2)))

	// log.Fatal(http.ListenAndServe(":8000", h3))
	// log.Fatal(http.ListenAndServe(":8000", h4))
}
