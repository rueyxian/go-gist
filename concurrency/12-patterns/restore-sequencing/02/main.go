package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// ============================================================
var rnd *rand.Rand

func init() {
	rnd = rand.New(rand.NewSource(time.Now().UnixNano()))
}

type Name struct {
	name     string
	count    int
	duration time.Duration
	wait     chan struct{}
}

// ============================================================

func main() {

	// names := []string{"Anna", "Bill", "Cass", "Dale", "Earl", "Finn", "Gabe"}
	names := []string{"Anna", "Bill", "Cass", "Dale"}
	// names := []string{"A     ", "BB    ", "CCC   ", "DDDD  ", "EEEEE ", "FFFFFF"}
	chs := []<-chan Name{}
	for _, name := range names {
		chs = append(chs, wait(name, 10))
	}

	ch := fanIn(chs)

	// ====================
	for v := range ch {
		fmt.Printf("%v  %v  %v\n", v.name, v.count, v.duration)
		v.wait <- struct{}{}
	}
	// ====================

}

// ============================================================
func wait(name string, count int) <-chan Name {
	ret := make(chan Name)
	go func() {
		defer close(ret)
		wait := make(chan struct{})
		for i := 0; i < count; i++ {
			duration := time.Duration(rnd.Intn(1000))
			message := Name{name: name, count: i, duration: duration, wait: wait}
			ret <- message
			time.Sleep(duration * time.Millisecond)
			<-wait
		}
	}()
	return ret
}

// ============================================================
func fanIn(chNames []<-chan Name) <-chan Name {
	var wg sync.WaitGroup
	wg.Add(len(chNames))
	ret := make(chan Name)

	output := func(ch <-chan Name) {
		defer wg.Done()
		for n := range ch {
			ret <- n
		}
	}

	go func() {
		defer close(ret)
		for _, ch := range chNames {
			go output(ch)
		}
		wg.Wait()
	}()

	return ret
}

// ============================================================

// ============================================================
