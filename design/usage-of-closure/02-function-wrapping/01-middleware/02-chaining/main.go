package main

import (
	"log"
	"math/rand"
	"net/http"
	"path"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {

	mux := http.NewServeMux()

	// mux.HandleFunc("/foo", timedMiddleware()(methodMiddleware("get")(getFoo)))
	// mux.HandleFunc("/bar", timedMiddleware()(methodMiddleware("get")(getBar)))

	// Be aware that the wrapping order might affect the outcome
	// In this case, say if we want time gets log if passing invalid method,
	// we want timedMiddleware goes first before methodMiddleware.
	mux.HandleFunc("/foo", chain(getFoo, methodMiddleware("GET"), timedMiddleware()))
	mux.HandleFunc("/bar", chain(getFoo, methodMiddleware("GET"), timedMiddleware()))

	server := http.Server{
		Addr:    "localhost:3000",
		Handler: mux,
	}

	server.ListenAndServe()

}

// ================================================================================

type middleware func(http.HandlerFunc) http.HandlerFunc

func timedMiddleware() middleware {
	return func(fn http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()
			fn(w, r)
			end := time.Now()
			url := path.Join(r.Host, r.URL.Path)
			log.Printf("%v %v : %v\n", r.Method, url, end.Sub(start))
		}
	}
}

func methodMiddleware(m string) middleware {
	return func(fn http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			if r.Method != m {
				http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
				return
			}
			fn(w, r)
		}
	}
}

func chain(fn http.HandlerFunc, mws ...middleware) http.HandlerFunc {
	for _, mw := range mws {
		fn = mw(fn)
	}
	return fn
}

// ================================================================================

func getFoo(w http.ResponseWriter, r *http.Request) {
	time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("get foo"))
}

func getBar(w http.ResponseWriter, r *http.Request) {
	time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("get bar"))
}
