package main

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
)

// ================================================================================

func main() {

	ts := httptest.NewUnstartedServer(http.HandlerFunc(helloHandler))
	ts.EnableHTTP2 = true
	ts.StartTLS()
	defer ts.Close()

	res, err := ts.Client().Get(ts.URL)
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
