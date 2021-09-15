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

	// The order of the middleware does matter. It might change the output of the system.

	// chain() function has changed, the order of the middleware ranging from 0 to n.

	// I personally prefer this than the previous one in this case:
	// handler is the last call stack (of course), and then index 0 middleware is the 2nd last,
	// index 1 middleware is the 3rd last, so on so forth.

	mux.handleFunc("/foo", chain(fooHandler, methodCheckMiddleware("GET"), errorMiddleware, timedMiddleware))
	mux.handleFunc("/bar", chain(barHandler, methodCheckMiddleware("GET"), errorMiddleware, timedMiddleware))

	// There is no hard rules on the order of middlewares,
	// it depends on what you are trying to achieve.
	//
	// For example, we want timedMiddleware calculates the whole process (including middlewares)
	// we have to put the timedMiddleware be top of the call stack;
	// If we want the timedMiddleware calculates only handler's process,

	http.ListenAndServe("localhost:8080", mux)

}

// ================================================================================
// middleware

type middleware func(handlerFunc) handlerFunc

func chain(fn handlerFunc, mws ...middleware) handlerFunc {

	// Changed the ranging order from 0 to n
	for _, mw := range mws {
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
