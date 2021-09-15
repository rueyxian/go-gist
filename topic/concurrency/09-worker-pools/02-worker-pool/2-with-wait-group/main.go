package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	workerCount := 2
	taskCount := 5

	var wg sync.WaitGroup

	chTasks := make(chan string, taskCount)
	chResults := make(chan string, taskCount)

	// spawning worker goroutines, those goroutines doing:
	// 1 - receiving tasks
	// 2 - process tasks and produce results
	// 3 - sending results
	for i := 0; i < workerCount; i++ {
		wg.Add(1)
		go worker(i, &wg, chTasks, chResults)
	}

	//sending tasks
	for i := 0; i < taskCount; i++ {
		task := string(i + 65)
		fmt.Printf("[main    ]   task   sending: %v \n", task)
		chTasks <- task
	}
	close(chTasks)

	wg.Wait()

	//receiving results
	for i := 0; i < taskCount; i++ {
		result := <-chResults
		fmt.Printf("[main    ] result receiving: %v \n", result)
	}

}

func worker(i int, wg *sync.WaitGroup, chTasks chan string, chResults chan string) {
	for task := range chTasks {
		fmt.Printf("[worker %v]   task receiving: %v \n", i, task)
		result := process(task)
		fmt.Printf("[worker %v] result   sending: %v \n", i, result)
		chResults <- result
	}
	wg.Done()
}

func process(task string) string {
	// simulate expensive process
	// time.Sleep(1000 * time.Millisecond)
	time.Sleep(time.Duration(rand.Intn(4000)) * time.Millisecond)

	return fmt.Sprintf("%s [processed]", task)
}
