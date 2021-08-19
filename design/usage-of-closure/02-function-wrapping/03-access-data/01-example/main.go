package main

import (
	"net/http"
	"time"
)

// ================================================================================

type DB struct{}

var db *DB

// Question: we need to access data from the database at the business logic,
// how do we do that?

// The easiest way is to declare database variable at the package level.
// However, simply declaring variable at the package level is a huge smell,
// it's almost a big no most of the time.

// ================================================================================

func main() {

	// ==============================

	mux := http.NewServeMux()
	mux.HandleFunc("/foo", fooHandler)
	mux.HandleFunc("/bar", barHandler)
	mux.HandleFunc("/baz", bazHandler)

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

func fooHandler(w http.ResponseWriter, r *http.Request) {
	data, _ := foo()
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}

func barHandler(w http.ResponseWriter, r *http.Request) {
	data, _ := bar()
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}

func bazHandler(w http.ResponseWriter, r *http.Request) {
	data, _ := baz()
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}

// ================================================================================
// business logic

func foo() ([]byte, error) {
	_ = db
	return []byte("foo"), nil
}

func bar() ([]byte, error) {
	_ = db
	return []byte("bar"), nil
}

func baz() ([]byte, error) {
	_ = db
	return []byte("baz"), nil
}
