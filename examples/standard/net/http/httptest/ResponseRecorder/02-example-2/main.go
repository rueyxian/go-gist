package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
)

func main() {

	{
		req := httptest.NewRequest("GET", "/it-doesn't-matter", nil)
		res := httptest.NewRecorder()

		dogoHandler(res, req)

		var v []map[string]interface{}

		if err := json.NewDecoder(res.Body).Decode(&v); err != nil {
			panic(err)
		}

		fmt.Println(v)

	}

}

// ================================================================================

func dogoHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(jsonString))
}

// ================================================================================

var jsonString = `
[
	{
		"Name": "Corgi",
		"Age": 3
	},
	{
		"Name": "Samoyad",
		"Age": 2
	}
]
`
