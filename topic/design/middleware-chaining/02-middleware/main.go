package main

import (
	"errors"
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

// ================================================================================
func main() {

	mux := newServeMux()
	// timedMiddleware wraps after errorMiddleware
	// that means that the timedMiddleware time not only the handler, also errorMiddleware
	mux.handleFunc("/foo", timedMiddleware(errorMiddleware(getFooHandler)))
	mux.handleFunc("/bar", timedMiddleware(errorMiddleware(getBarHandler)))
	http.ListenAndServe("localhost:8080", mux)

}

// ================================================================================
// middleware

func timedMiddleware(fn handlerFunc) handlerFunc {
	return func(w http.ResponseWriter, r *http.Request) error {
		start := time.Now()
		err := fn(w, r)
		fmt.Printf("%s %s : %v\n", r.Method, r.URL, time.Now().Sub(start))
		return err
	}
}

func errorMiddleware(fn handlerFunc) handlerFunc {
	return func(w http.ResponseWriter, r *http.Request) error {
		if err := fn(w, r); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return nil // error has been handled, return nil
	}
}

// ================================================================================
// mux / router

type handlerFunc func(http.ResponseWriter, *http.Request) error

type serveMux struct {
	mux *http.ServeMux
}

func newServeMux() *serveMux {
	return &serveMux{mux: http.NewServeMux()}
}

func (m *serveMux) handleFunc(pattern string, handlerFn handlerFunc) {
	fn := func(w http.ResponseWriter, r *http.Request) {
		if err := handlerFn(w, r); err != nil {
			fmt.Printf("unhandle error: %s\n", err)
		}
	}
	m.mux.HandleFunc(pattern, fn)
}

func (m *serveMux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	m.mux.ServeHTTP(w, r)
}

// ================================================================================
// handlers

func getFooHandler(w http.ResponseWriter, r *http.Request) error {
	time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
	if r.Method != "GET" {
		return newWebError(errors.New("error: wrong request method"), http.StatusBadRequest)
	}
	if rand.Intn(2) == 0 {
		return newWebError(errors.New("error: db query error"), http.StatusInternalServerError)
	}
	w.Write([]byte("foo"))
	return nil
}

func getBarHandler(w http.ResponseWriter, r *http.Request) error {
	time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
	if r.Method != "GET" {
		return newWebError(errors.New("error: wrong request method"), http.StatusBadRequest)
	}
	if rand.Intn(2) == 0 {
		return newWebError(errors.New("error: db query error"), http.StatusInternalServerError)
	}
	w.Write([]byte("bar"))
	return nil
}

// ================================================================================
// web error

type webError struct {
	err    error
	status int
}

func newWebError(err error, status int) *webError {
	return &webError{err: err, status: status}
}

func (e *webError) Error() string {
	return e.err.Error()
}

// ================================================================================
