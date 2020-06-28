package main

import (
	"fmt"
	"math/rand"
	"time"
)

type data struct {
	task   int
	result []int
}

func main() {

	rand.Seed(time.Now().UnixNano())
	workerCount := 3
	taskCount := 10

	// tasks := make(chan int, workerCount+1)
	// results := make(chan data, workerCount+1)
	tasks := make(chan int, workerCount)
	results := make(chan data, workerCount)

	// spawning worker goroutines, those goroutines doing:
	// 1 - receiving tasks
	// 2 - process tasks and produce results
	// 3 - sending results
	for i := 0; i < workerCount; i++ {
		go fibonacciWorker(i, tasks, results)
	}

	//sending tasks
	for _, task := range rand.Perm(20)[taskCount:] {
		tasks <- task
		fmt.Printf("[main    ]   task   sending: %v \n", task)
	}
	close(tasks)

	//receiving results
	for i := 0; i < taskCount; i++ {
		fmt.Printf("[main    ] result receiving: %+v \n", <-results)
	}

}

func fibonacciWorker(workerNum int, tasks <-chan int, results chan<- data) {
	for t := range tasks {
		fmt.Printf("[worker %v]   task receiving: %v \n", workerNum, t)
		result := data{t, fibonacci(t)}
		fmt.Printf("[worker %v] result   sending: %+v\n", workerNum, result)
		results <- result
	}
}

func fibonacci(n int) []int {
	// time.Sleep(time.Duration(rand.Intn(4000)) * time.Millisecond) // simulate blocking heavy process
	output := make([]int, n)
	a, b := 1, 0
	for i := 0; i < n; i++ {
		a, b = b, a+b
		output[i] = a
	}
	return output
}
