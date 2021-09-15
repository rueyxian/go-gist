package main

import (
	"flag"
	"log"
	"math/rand"
	"os"
	"time"
)

func main() {

	foo := flag.String("foo", "default", "help foo")
	bar := flag.Int("bar", 4896, "help bar")

	flag.Parse()

	rand.Seed(time.Now().UnixNano())
	stdout := log.New(os.Stdout, "", log.LstdFlags)
	stderr := log.New(os.Stderr, "", log.LstdFlags)

	if rand.Intn(2) == 0 {
		stderr.Print("stderr")
		return
	}

	stdout.Print(*foo, " ", *bar)
}
