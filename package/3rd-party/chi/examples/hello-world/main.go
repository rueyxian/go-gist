package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
)

func main() {
	r := chi.NewRouter()

	r.Get("/", indexHandler)
	r.Get("/hello", helloHandler)

	http.ListenAndServe(":3000", r)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "index")
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "hello")
}
