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

	// mux := newServeMux()
	// mux.handleFunc("/foo", chain(fooHandler, methodCheckMiddleware("GET"), errorMiddleware, timedMiddleware))
	// mux.handleFunc("/bar", chain(barHandler, methodCheckMiddleware("GET"), errorMiddleware, timedMiddleware))

	// Imagine that we have lots of handlerFunc,
	// it would be cumbersome to do all these chaining thing.
	// So it will be nice if we store the middlewares as a field
	// of our custom serveMux.
	// Do the chaining inside mux.handleFunc method instead of outside
	mux := newServeMux(timedMiddleware, errorMiddleware, methodCheckMiddleware("GET"))
	// noted that the order of has changed from index n to index 0,
	// and the two parameters of chain function has swapped.

	mux.handleFunc("/foo", fooHandler)
	mux.handleFunc("/bar", barHandler)

	http.ListenAndServe("localhost:8080", mux)

}

// ================================================================================
// middleware

type middleware func(handlerFunc) handlerFunc

func chain(mws []middleware, fn handlerFunc) handlerFunc {

	for i := len(mws) - 1; i >= 0; i-- {
		mw := mws[i]
		fn = mw(fn)
	}
	return fn
}

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
			if webErr, ok := err.(*webError); ok {
				http.Error(w, err.Error(), webErr.status)
				return nil
			}
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return nil // error has been handled, return nil
	}
}

func methodCheckMiddleware(method string) middleware {
	return func(fn handlerFunc) handlerFunc {
		return func(w http.ResponseWriter, r *http.Request) error {
			if r.Method != method {
				return newWebError(errors.New("error: wrong request method"), http.StatusBadRequest)
			}
			return fn(w, r)
		}
	}
}

// ================================================================================
// mux / router

type handlerFunc func(http.ResponseWriter, *http.Request) error

type serveMux struct {
	mux *http.ServeMux
	mws []middleware
}

func newServeMux(mws ...middleware) *serveMux {
	return &serveMux{mux: http.NewServeMux(), mws: mws}
}

func (m *serveMux) handleFunc(pattern string, handlerFn handlerFunc) {
	fn := func(w http.ResponseWriter, r *http.Request) {

		handlerFn = chain(m.mws, handlerFn)

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

func fooHandler(w http.ResponseWriter, r *http.Request) error {
	time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)

	if rand.Intn(2) == 0 {
		return newWebError(errors.New("error: db query error"), http.StatusInternalServerError)
	}
	w.Write([]byte("foo"))
	return nil
}

func barHandler(w http.ResponseWriter, r *http.Request) error {
	time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)

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
