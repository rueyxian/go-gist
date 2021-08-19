package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type response1 struct {
	Page   int
	Fruits []string
}

type response2 struct {
	Page   int      `json:"page"`
	Fruits []string `json:"fruits"`
}

func main() {

	// bolB, _ := json.Marshal(true)
	// fmt.Println(string(bolB))

	// intB, _ := json.Marshal(1)
	// fmt.Println(string(intB))

	// fltB, _ := json.Marshal(2.34)
	// fmt.Println(string(fltB))

	// strB, _ := json.Marshal("gopher")
	// fmt.Println(string(strB))

	{
		v := []string{"apple", "peach", "pear"}
		b, _ := json.MarshalIndent(v, "", "  ")
		fmt.Println(string(b))
		fmt.Println()
	}

	{
		v := map[string]int{"apple": 5, "lettuce": 7}
		b, _ := json.MarshalIndent(v, "", "  ")
		fmt.Println(string(b))
		fmt.Println()
	}

	{
		v := response1{
			Page:   1,
			Fruits: []string{"apple", "peach", "pear"}}
		b, _ := json.MarshalIndent(&v, "", "  ")
		fmt.Println(string(b))
		fmt.Println()
	}

	{
		v := &response2{
			Page:   1,
			Fruits: []string{"apple", "peach", "pear"}}
		b, _ := json.MarshalIndent(&v, "", "  ")
		fmt.Println(string(b))
		fmt.Println()
	}

	// ==============================
	fmt.Println()
	fmt.Println("============================================================")
	fmt.Println()
	// ==============================

	{
		b := []byte(`{"num":6.13,"strs":["a","b"]}`)
		var v map[string]interface{}
		json.Unmarshal(b, &v)
		fmt.Println(v)
		fmt.Println()
	}

	{
		b := []byte(`{"page": 1, "fruits": ["apple", "peach"]}`)
		v := response2{}
		json.Unmarshal(b, &v)
		fmt.Println(v)
		fmt.Println()
	}

	// ==============================
	fmt.Println()
	fmt.Println("============================================================")
	fmt.Println()
	// ==============================

	enc := json.NewEncoder(os.Stdout)
	d := map[string]int{"apple": 5, "lettuce": 7}
	enc.Encode(d)

	// ==============================
	fmt.Println()
	fmt.Println("============================================================")
	fmt.Println()
	// ==============================
}
