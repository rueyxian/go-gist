package main

import (
	"bufio"
	"fmt"
	"net/http"
)

func main() {

	res, err := http.Get("https://golang.org/pkg/net/http/")
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	fmt.Printf("%T \n", res)

	fmt.Printf(" status: %v \n", res.Status)
	// fmt.Printf("body: %v \n", resp.Body)

	scanner := bufio.NewScanner(res.Body)
	for i := 0; scanner.Scan() && i < 5; i++ {
		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}
}
