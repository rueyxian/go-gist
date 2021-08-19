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

// ================================================================================

type HandlerFunc func(w http.ResponseWriter, r *http.Request) error

// ================================================================================

func main() {

	mux := http.NewServeMux()

	handleFunc(mux, "/foo", fooHandler)
	handleFunc(mux, "/bar", barHandler)
	handleFunc(mux, "/baz", bazHandler)

	server := http.Server{
		Addr:         "localhost:3000",
		Handler:      mux,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
	}

	server.ListenAndServe()
}

// ================================================================================
// Instead of making an adapter function to transform HandlerFunc signature,
// we can improve the adapter by accepting *http.ServeMux as an argument (and also path pattern)
//
// But we can do even better than this, as you can see
// handleFunc(mux, "/foo", fooHandler)
// handleFunc(mux, "/bar", barHandler)
// handleFunc(mux, "/baz", bazHandler)
// The argument mux (*http.ServeMux) is repetitive
//
// Note: This improvement (and subsequence) is nothing to do with closure though.

func handleFunc(mux *http.ServeMux, pattern string, handlerFn HandlerFunc) {
	fn := func(w http.ResponseWriter, r *http.Request) {
		if err := handlerFn(w, r); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			log.Printf("ERROR : %v\n", err)
		}
	}
	mux.HandleFunc(pattern, fn)
}

// ================================================================================
// business logic

func fooHandler(w http.ResponseWriter, r *http.Request) error {
	data, err := foo()
	if err != nil {
		return err
	}
	w.WriteHeader(http.StatusOK)
	w.Write(data)
	return nil
}

func barHandler(w http.ResponseWriter, r *http.Request) error {
	data, err := bar()
	if err != nil {
		return err
	}
	w.WriteHeader(http.StatusOK)
	w.Write(data)
	return nil
}

func bazHandler(w http.ResponseWriter, r *http.Request) error {
	data, err := baz()
	if err != nil {
		return err
	}
	w.WriteHeader(http.StatusOK)
	w.Write(data)
	return nil
}

// ================================================================================

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
