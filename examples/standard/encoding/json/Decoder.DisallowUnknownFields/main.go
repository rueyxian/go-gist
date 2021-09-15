package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Dog struct {
	Breed string `json:"breed"`
	Age   int    `json:"age"`
}

func main() {
	f, err := os.Open("data.json")
	if err != nil {
		panic(err)
	}

	decoder := json.NewDecoder(f)
	decoder.DisallowUnknownFields()

	var v Dog
	if err := decoder.Decode(&v); err != nil {
		panic(err)
	}

	fmt.Println(v)

}
