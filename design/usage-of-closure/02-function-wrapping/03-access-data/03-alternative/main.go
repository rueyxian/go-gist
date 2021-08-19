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

	fb := foobar{db: db}
	// ==============================

	mux := http.NewServeMux()
	mux.HandleFunc("/foo", fb.fooHandler)
	mux.HandleFunc("/bar", fb.barHandler)
	mux.HandleFunc("/baz", fb.bazHandler)

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
// Again, this is nothing to do with closure.

type foobar struct {
	db *DB
}

func (fb *foobar) fooHandler(w http.ResponseWriter, r *http.Request) {
	data, _ := foo(fb.db)
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}

func (fb *foobar) barHandler(w http.ResponseWriter, r *http.Request) {
	data, _ := bar(fb.db)
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}

func (fb *foobar) bazHandler(w http.ResponseWriter, r *http.Request) {
	data, _ := baz(fb.db)
	w.WriteHeader(http.StatusOK)
	w.Write(data)
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
