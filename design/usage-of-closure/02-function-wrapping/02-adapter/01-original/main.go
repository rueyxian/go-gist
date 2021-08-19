package main

import (
	"errors"
	"log"
	"math/rand"
	"net/http"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {

	mux := http.NewServeMux()

	mux.HandleFunc("/foo", fooHandler)
	mux.HandleFunc("/bar", barHandler)
	mux.HandleFunc("/baz", bazHandler)

	server := http.Server{
		Addr:         "localhost:3000",
		Handler:      mux,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
	}

	server.ListenAndServe()

}

// ================================================================================
// How do we centralize the error handling of these handler?
// The answer is to return error interface.
// However, by doing that, these handler will no longer implement http.HanderFunc

// It seems superfluous to to that, but in production level, the code are more complex.
// It always nice to be centralized.

func fooHandler(w http.ResponseWriter, r *http.Request) {
	data, err := foo()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Printf("ERROR : %v\n", err)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}

func barHandler(w http.ResponseWriter, r *http.Request) {
	data, err := bar()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Printf("ERROR : %v\n", err)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}

func bazHandler(w http.ResponseWriter, r *http.Request) {
	data, err := baz()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Printf("ERROR : %v\n", err)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}

// ================================================================================
// business logic

func foo() ([]byte, error) {
	if rand.Intn(2) == 0 {
		return []byte("foo"), nil
	}
	return nil, errors.New("error: foo")
}

func bar() ([]byte, error) {
	if rand.Intn(2) == 0 {
		return []byte("bar"), nil
	}
	return nil, errors.New("error: bar")
}

func baz() ([]byte, error) {
	if rand.Intn(2) == 0 {
		return []byte("baz"), nil
	}
	return nil, errors.New("error: baz")
}
