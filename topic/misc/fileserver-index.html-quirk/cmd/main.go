package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

// ================================================================================

func newRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/", indexHandler).Methods("GET")
	r.HandleFunc("/hello", helloHandler).Methods("GET")
	r.PathPrefix("/static/").Handler(staticFileHandler()).Methods("GET")
	return r
}

// func newRouter() *http.ServeMux {
//   r := http.NewServeMux()
//   r.HandleFunc("/", indexHandler)
//   r.HandleFunc("/hello", helloHandler)
//   r.Handle("/static/", staticFileHandler())
//   return r
// }

// ================================================================================

func indexHandler(w http.ResponseWriter, r *http.Request) {
	b, err := os.ReadFile("../resources/index/index.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	fmt.Fprintf(w, string(b))
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	b, err := os.ReadFile("../resources/hello/hello.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	fmt.Fprintf(w, string(b))
}

func staticFileHandler() http.Handler {
	dir := http.Dir("../")
	return http.StripPrefix("/static/", http.FileServer(dir))
}

// ================================================================================

func main() {
	r := newRouter()
	http.ListenAndServe(":8080", r)
}

// ================================================================================
