package main

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
)

func main() {

	{
		req := httptest.NewRequest("GET", "/it-doesn't-matter", strings.NewReader("corgi"))
		w := httptest.NewRecorder()

		helloHandler(w, req)

		res := w.Result()
		body, _ := io.ReadAll(res.Body)
		fmt.Println(string(body))
	}

	{
		req := httptest.NewRequest("GET", "/it-doesn't-matter", strings.NewReader("samoyad"))
		res := httptest.NewRecorder()

		byeHandler(res, req)

		body, _ := io.ReadAll(res.Body)
		fmt.Println(string(body))
	}

}

// ================================================================================

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
