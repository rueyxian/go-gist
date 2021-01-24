package main

import (
	"fmt"
	"io/ioutil"
	"sandbox/go-jottings/design/error-handling/03-types-as-context/dsv"
)

type item struct {
	id    int
	name  string
	price float64
}

func main() {

	var items []item

	// var test item

	b, err := ioutil.ReadFile("list_1.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	if err := dsv.Unmarshal(b, items); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Peace out~")

}
