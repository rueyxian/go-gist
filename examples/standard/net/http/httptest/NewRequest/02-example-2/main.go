package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
)

// ================================================================================

func main() {

	{
		req := httptest.NewRequest("GET", "/dogo", nil)
		res := httptest.NewRecorder()

		routers().ServeHTTP(res, req)

		var v []map[string]interface{}

		if err := json.NewDecoder(res.Body).Decode(&v); err != nil {
			panic(err)
		}
		fmt.Println(v)
	}

}

// ================================================================================

func routers() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/dogo", dogoHandler)
	return mux
}

func dogoHandler(w http.ResponseWriter, r *http.Request) {
	if _, err := w.Write([]byte(jsonString)); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// ================================================================================

var jsonString = `
[
	{
		"Name": "Corgi",
		"Age": 2
	},
	{
		"Name": "Samoyad",
		"Age": 3
	},
	{
		"Name": "Shiba Inu",
		"Age": 1
	}
]
`
