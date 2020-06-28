package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	workerCount := 2
	taskCount := 5

	tasks := make(chan string, taskCount)
	results := make(chan string, taskCount)

	// spawning worker goroutines, those goroutines doing:
	// 1 - receiving tasks
	// 2 - process tasks and produce results
	// 3 - sending results
	for i := 0; i < workerCount; i++ {
		go worker(i, tasks, results)
	}

	//sending tasks
	for i := 0; i < taskCount; i++ {
		task := string(i + 65)
		fmt.Printf("[main    ]   task   sending: %v \n", task)
		tasks <- task
	}
	close(tasks)

	//receiving results
	for i := 0; i < taskCount; i++ {
		result := <-results
		fmt.Printf("[main    ] result receiving: %v \n", result)
	}

}

func worker(i int, tasks chan string, results chan string) {
	for task := range tasks {
		fmt.Printf("[worker %v]   task receiving: %v \n", i, task)
		result := process(task)
		fmt.Printf("[worker %v] result   sending: %v \n", i, result)
		results <- result
	}
}

func process(task string) string {
	// simulate blocking process
	// time.Sleep(1000 * time.Millisecond)
	time.Sleep(time.Duration(rand.Intn(4000)) * time.Millisecond)

	return fmt.Sprintf("%s [processed]", task)
}
