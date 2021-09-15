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
// We have made serveral changes here:
// 1) Define our own mux struct that wraps *http.ServerMux
// 2) Make previously defined handleFunc (the adapter) as method of our mux struct
// 3) mux struct implements http.Handler – ServerHTTP(w http.ReseWriter, r *http.Request) signature.

type mux struct {
	mux *http.ServeMux
}

func newMux() *mux {
	return &mux{http.NewServeMux()}
}

func (m *mux) handleFunc(pattern string, handlerFn HandlerFunc) {
	fn := func(w http.ResponseWriter, r *http.Request) {
		if err := handlerFn(w, r); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			log.Printf("ERROR : %v\n", err)
		}
	}
	m.mux.HandleFunc(pattern, fn)
}

func (m *mux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	m.mux.ServeHTTP(w, r)
}

// ================================================================================

func main() {

	muxFn := func() http.Handler {
		mux := newMux()
		mux.handleFunc("/foo", fooHandler)
		mux.handleFunc("/bar", barHandler)
		mux.handleFunc("/baz", bazHandler)
		return mux
	}

	// ==============================

	server := http.Server{
		Addr:         "localhost:3000",
		Handler:      muxFn(),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
	}

	server.ListenAndServe()
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
