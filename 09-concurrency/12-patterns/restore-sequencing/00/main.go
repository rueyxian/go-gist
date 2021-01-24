package main

import (
	"fmt"
	"math/rand"
	"time"
)

var rnd *rand.Rand

func init() {
	rnd = rand.New(rand.NewSource(time.Now().UnixNano()))
}

// ================================================================================
func main() {
	ch1 := boring("anna")
	ch2 := boring("bill")
	ch := fanIn(ch1, ch2)

	for v := range ch {
		fmt.Println(v)
	}
}

// ================================================================================
func boring(msg string) <-chan string {
	ret := make(chan string)
	go func() {
		for i := 0; ; i++ {
			ret <- fmt.Sprintf("%s: %d", msg, i)
			duration := time.Duration(rnd.Intn(1e3))
			time.Sleep(duration * time.Millisecond)
		}
	}()
	return ret
}

// ================================================================================
func fanIn(chs ...<-chan string) <-chan string {
	ret := make(chan string)

	output := func(ch <-chan string) {
		for {
			ret <- <-ch
		}
	}

	go func() {
		for _, ch := range chs {
			go output(ch)
		}
	}()

	// go func() {
	//   for {
	//     ret <- <-ch1
	//   }
	// }()

	// go func() {
	//   for {
	//     ret <- <-ch2
	//   }
	// }()

	return ret
}
