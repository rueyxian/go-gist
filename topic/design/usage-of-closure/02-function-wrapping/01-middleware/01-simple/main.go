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
	mux.HandleFunc("/foo", timed(foo))
	mux.HandleFunc("/bar", timed(bar))

	server := http.Server{
		Addr:    "localhost:3000",
		Handler: mux,
	}

	server.ListenAndServe()

}

// ================================================================================

func timed(fn func(w http.ResponseWriter, r *http.Request)) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		fn(w, r)
		end := time.Now()
		url := path.Join(r.Host, r.URL.Path)
		log.Printf("%v %v : %v\n", r.Method, url, end.Sub(start))
	}
}

// ================================================================================

func foo(w http.ResponseWriter, r *http.Request) {
	time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("get foo"))
}

func bar(w http.ResponseWriter, r *http.Request) {
	time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("get bar"))
}
