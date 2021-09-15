package main

import (
	"log"
	"math/rand"
	"os"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	stdout := log.New(os.Stdout, "", log.LstdFlags)
	stderr := log.New(os.Stderr, "", log.LstdFlags)

	time.Sleep(time.Second)

	if rand.Intn(2) == 0 {
		stderr.Print("stderr")
		return
	}

	stdout.Print("stdout")
}
