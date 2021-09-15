package main

import (
	"fmt"
	"io"
	"net/http"
)

// ================================================================================

func main() {

	server := http.Server{
		Addr:    "localhost:3000",
		Handler: routers(),
	}

	go func() {
		server.ListenAndServe()
	}()

	// ==============================

	{
		// http.Get calling package level variable's method – DefaultClient
		res, err := http.Get("http://localhost:3000/foo")
		if err != nil {
			panic(err)
		}

		b, err := io.ReadAll(res.Body)
		defer res.Body.Close()
		if err != nil {
			panic(err)
		}

		fmt.Println(string(b))
	}

	// ==============================

	{
		// http.Get calling package level variable's method – DefaultClient
		res, err := http.Get("http://localhost:3000/bar")
		if err != nil {
			panic(err)
		}
		b, err := io.ReadAll(res.Body)
		defer res.Body.Close()
		if err != nil {
			panic(err)
		}
		fmt.Println(string(b))
	}

}

// ================================================================================

func routers() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/foo", fooHandler)
	mux.HandleFunc("/bar", barHandler)
	return mux
}

func fooHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("foo"))
}

func barHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("bar"))
}

// ================================================================================

type item struct {
	Name     string
	Quantity int
}
