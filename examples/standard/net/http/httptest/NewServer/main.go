package main

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
)

// ================================================================================

func main() {

	ts := httptest.NewServer(http.HandlerFunc(helloHandler))
	defer ts.Close()

	res, err := http.Get(ts.URL)
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

// ================================================================================

func helloHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello gopher"))
}
