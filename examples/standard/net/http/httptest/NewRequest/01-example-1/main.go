package main

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
)

// ================================================================================

func main() {

	{
		req := httptest.NewRequest("POST", "/hello", strings.NewReader("gopher <3"))
		recorder := httptest.NewRecorder()

		r := routers()
		r.ServeHTTP(recorder, req)

		res := recorder.Result()
		body, _ := io.ReadAll(res.Body)
		fmt.Println(string(body))
	}

	// ==============================
	fmt.Println()
	// ==============================

	{
		req := httptest.NewRequest("POST", "/bye", strings.NewReader("gopher ^.^"))
		res := httptest.NewRecorder()

		r := routers()
		r.ServeHTTP(res, req)

		body, _ := io.ReadAll(res.Body)
		fmt.Println(string(body))
	}

}

// ================================================================================

func routers() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/hello", helloHandler)
	mux.HandleFunc("/bye", byeHandler)
	return mux
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	b, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	fmt.Fprintf(w, "hello %s", string(b))
}

func byeHandler(w http.ResponseWriter, r *http.Request) {
	b, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	fmt.Fprintf(w, "bye %s", string(b))
}

// ================================================================================
