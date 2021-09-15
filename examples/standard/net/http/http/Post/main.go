package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
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
		url := "http://localhost:3000/hello"
		content := "text/html; charset=UTF-8"
		data := strings.NewReader("gopher o.o")

		// http.Post calling package level variable's method – DefaultClient
		res, err := http.Post(url, content, data)
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
	mux.HandleFunc("/hello", helloHandler)
	// mux.HandleFunc("/foo", fooHandler)
	// mux.HandleFunc("/bar", barHandler)
	return mux
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	b, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	fmt.Fprintf(w, "hello %s\n", string(b))
}

// func fooHandler(w http.ResponseWriter, r *http.Request) {
//   w.Write([]byte("foo"))
// }

// func barHandler(w http.ResponseWriter, r *http.Request) {
//   w.Write([]byte("bar"))
// }

// ================================================================================
