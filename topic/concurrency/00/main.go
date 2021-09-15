package main

import (
	"math/rand"
	"time"
)

var rnd *rand.Rand

func init() {
	rnd = rand.New(rand.NewSource(time.Now().UnixNano()))
}

func main() {

	operation()

}

func operation() {

	ch := make(chan time.Duration)
	done := make(chan struct{})

	ping(done, ch)

}

func ping(done chan struct{}, ch chan time.Duration) {
	for {
		duration := time.Duration(rnd.Intn(1000) * time.Millisecond)		
		time.Sleep(duration)
		select{
			ch <- duration:

		}
	}
}
