package main

import (
	"log"
	"math/rand"
	"os"
	"time"
)

// ================================================================================

func init() {
	rand.Seed(time.Now().UnixNano())
}

// ================================================================================

func main() {
	rand.Seed(time.Now().UnixNano())
	stdout := log.New(os.Stdout, "", 0)
	stderr := log.New(os.Stderr, "", 0)

	p := progress(100)
	for range p {

		if rand.Intn(100) == 0 {
			stderr.Print("unable to finish")
			return
		}

		stdout.Print()
	}

	// stdout.Print("finished")
}

// ================================================================================
// simulate progress

func progress(n int) <-chan struct{} {
	out := make(chan struct{})
	go func() {
		defer close(out)
		for i := 1; i <= n; i++ {
			doTask()
			// out <- fmt.Sprintf("%d/%d", i, n)
			out <- struct{}{}
		}
	}()
	return out
}

// func progress(n int) <-chan string {
//   out := make(chan string)
//   go func() {
//     defer close(out)
//     for i := 1; i <= n; i++ {
//       doTask()
//       // out <- fmt.Sprintf("%d/%d", i, n)
//       out <- ""
//     }
//   }()
//   return out
// }

func doTask() {
	// time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
	time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
}
