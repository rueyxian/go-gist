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
	// We can also delegate request method checking to middleware too
	mux.handleFunc("/foo", chain(fooHandler, timedMiddleware, errorMiddleware, methodCheckMiddleware("GET")))
	mux.handleFunc("/bar", chain(barHandler, timedMiddleware, errorMiddleware, methodCheckMiddleware("GET")))

	http.ListenAndServe("localhost:8080", mux)

}

// ================================================================================
// middleware

type middleware func(handlerFunc) handlerFunc

func chain(fn handlerFunc, mws ...middleware) handlerFunc {
	// Ranging from index N to index 0.
	// Meaning, index 0 middleware will be the most outer of of call stack,
	// while index N will be the most inner of the call stack
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
				http.Error(w, webErr.err.Error(), webErr.status)
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
				// if wrong request, pop up right away of the call stack
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

func fooHandler(w http.ResponseWriter, r *http.Request) error {
	time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
	// if r.Method != "GET" {
	//   return newWebError(errors.New("error: wrong request method"), http.StatusBadRequest)
	// }
	if rand.Intn(2) == 0 {
		return newWebError(errors.New("error: db query error"), http.StatusInternalServerError)
	}
	w.Write([]byte("foo"))
	return nil
}

func barHandler(w http.ResponseWriter, r *http.Request) error {
	time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
	// if r.Method != "GET" {
	//   return newWebError(errors.New("error: wrong request method"), http.StatusBadRequest)
	// }
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
