package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/gorilla/schema"
)

type Person struct {
	Name  string
	Level int
}

// ================================================================================

// func newRouter() *mux.Router {
//   r := mux.NewRouter()
//   r.HandleFunc("/hello", handler).Methods("POST")
//   return r
// }

func main() {

	// mux := http.NewServeMux()
	// mux.HandleFunc("/hello", helloHandler)

	// http.ListenAndServe(":8080", mux)

	// ==============================

	r := mux.NewRouter()
	r.HandleFunc("/hello", helloHandler).Methods("POST")
	r.HandleFunc("/bye", byeHandler).Methods("GET")
	r.HandleFunc("/login", loginHandler).Methods("POST")

	http.ListenAndServe(":8080", r)

	// ==============================

}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	var err error
	var p Person

	err = json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	fmt.Fprintf(w, "Hello : %v", p)
}

func byeHandler(w http.ResponseWriter, r *http.Request) {
	var err error
	var p Person

	err = schema.NewDecoder().Decode(&p, r.URL.Query())
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	fmt.Fprintf(w, "Bye : %v ", p)
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	// var err error
	// h := r.Header.Get("x-auth-token")
	if r.Header.Get("x-auth-token") == "12345" {
		fmt.Fprintf(w, "login")
	} else {
		http.Error(w, "%v", http.StatusNonAuthoritativeInfo)
		return
	}

}
