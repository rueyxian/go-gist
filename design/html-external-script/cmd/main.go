package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

// ================================================================================

func main() {
	r := newRouter()
	http.ListenAndServe(":8080", r)
}

// ================================================================================

func newRouter() *mux.Router {
	r := mux.NewRouter()
	r.PathPrefix("/assets/").Handler(staticFileHandler()).Methods("GET")
	r.HandleFunc("/", indexHandler).Methods("GET")
	r.HandleFunc("/hello", helloHandler).Methods("GET")
	return r
}

func staticFileHandler() http.Handler {
	dir := http.Dir("../")
	return http.StripPrefix("/assets/", http.FileServer(dir))
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	b, err := os.ReadFile("../assets/index/index.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	fmt.Fprintf(w, string(b))
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	b, err := os.ReadFile("../assets/hello/hello.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	fmt.Fprintf(w, string(b))
}
