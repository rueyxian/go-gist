package main

import (
	"net/http"
	"time"
)

// ================================================================================

type DB struct{}

// ================================================================================

func main() {

	var db *DB

	// ==============================

	mux := http.NewServeMux()
	mux.HandleFunc("/foo", fooHandler(db))
	mux.HandleFunc("/bar", barHandler(db))
	mux.HandleFunc("/baz", bazHandler(db))

	// ==============================

	server := http.Server{
		Addr:         "localhost:3000",
		Handler:      mux,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
	}

	server.ListenAndServe()
}

// ================================================================================
// One way to work with it is to use closure technique.
// But, in this case, closure actually not the best solution.
// Given that when we are working on web service's handler,
// often time, we need a lot of http.HandlerFunc.
//
// Nothing wrong with closure, it's actually an expediant
// technique in certain situation.

func fooHandler(db *DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data, _ := foo(db)
		w.WriteHeader(http.StatusOK)
		w.Write(data)
	}
}

func barHandler(db *DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data, _ := bar(db)
		w.WriteHeader(http.StatusOK)
		w.Write(data)
	}
}

func bazHandler(db *DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data, _ := baz(db)
		w.WriteHeader(http.StatusOK)
		w.Write(data)
	}
}

// ================================================================================
// business logic

func foo(db *DB) ([]byte, error) {
	_ = db
	return []byte("foo"), nil
}

func bar(db *DB) ([]byte, error) {
	_ = db
	return []byte("bar"), nil
}

func baz(db *DB) ([]byte, error) {
	_ = db
	return []byte("baz"), nil
}
