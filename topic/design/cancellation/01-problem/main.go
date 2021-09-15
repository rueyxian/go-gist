package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

// ================================================================================

// When we send a request and cancel at the midst of the operation,
// the process continues until it completes.
// This definitely not ideal in that request had already cancelled,
// the process should be also cancelled

func main() {

	addr := "localhost:8000"

	http.HandleFunc("/", simpleHandler)

	fmt.Printf("listening %s\n", addr)
	log.Fatal(http.ListenAndServe(addr, nil))

}

// ================================================================================

func simpleHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("start")
	defer log.Println("end\n")

	operation()
}

// ================================================================================

func operation() {
	for i, n := 0, 10; i < n; i++ {
		time.Sleep(time.Second)
		log.Printf("progress %d/%d\n", i+1, n)
	}
}

// ================================================================================
