package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

func main() {

	r := strings.NewReader("Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua")

	b, err := ioutil.ReadAll(r)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%v", b)

	fmt.Println()
	fmt.Println()

	fmt.Printf("%s", b)

}
