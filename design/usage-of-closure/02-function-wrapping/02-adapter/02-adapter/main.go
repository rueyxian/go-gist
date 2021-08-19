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

	mux.HandleFunc("/foo", adapter(fooHandler))
	mux.HandleFunc("/bar", adapter(barHandler))
	mux.HandleFunc("/baz", adapter(bazHandler))

	server := http.Server{
		Addr:         "localhost:3000",
		Handler:      mux,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
	}

	server.ListenAndServe()
}

// ================================================================================
// The adapter function act as an adapter between our custom HandlerFunc and http.HandlerFunc.
// This a very primitive adapter, but we can do more than that.
// More than just the transformation of function signature.
//
// And also the code here:
// mux.HandleFunc("/foo", adapter(fooHandler))
// mux.HandleFunc("/bar", adapter(barHandler))
// mux.HandleFunc("/baz", adapter(bazHandler))
// is somehow awkward.

func adapter(handlerFn HandlerFunc) http.HandlerFunc {
	fn := func(w http.ResponseWriter, r *http.Request) {
		if err := handlerFn(w, r); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			log.Printf("ERROR : %v\n", err)
		}
	}
	return http.HandlerFunc(fn)
}

// ================================================================================

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
