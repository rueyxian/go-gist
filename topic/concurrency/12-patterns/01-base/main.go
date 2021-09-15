package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var rnd *rand.Rand

func init() {
	rnd = rand.New(rand.NewSource(time.Now().UnixNano()))
}

func main() {
	waitForTask()
	// waitForResult()

}

func waitForTask() {

	var wg sync.WaitGroup
	ch := make(chan string)

	wg.Add(1)

	// ====================
	go func() {
		defer wg.Done()
		task := <-ch
		fmt.Printf("worker  [recv]   data: %v\n", task)
	}()

	// ====================
	duration := time.Duration(rnd.Intn(500)) * time.Millisecond
	time.Sleep(duration)

	ch <- "task"
	fmt.Printf("manager [send]   duration: %v\n", duration)

	// ====================
	wg.Wait()
	fmt.Println("==============================")

}

func waitForResult() {

	var wg sync.WaitGroup
	ch := make(chan string)

	wg.Add(1)

	// ====================
	go func() {
		defer wg.Done()

		duration := time.Duration(rnd.Intn(500)) * time.Millisecond
		time.Sleep(duration)

		ch <- "result"
		fmt.Printf("worker  [send]   duration: %v\n", duration)
	}()

	// ====================
	result := <-ch
	fmt.Printf("manager [recv]   data: %v\n", result)

	// ====================
	wg.Wait()
	fmt.Println("==============================")

}
