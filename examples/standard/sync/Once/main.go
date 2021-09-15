package main

import (
	"fmt"
	"sync"
)

func main() {

	var once sync.Once
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		onceOperation := func() {
			fmt.Printf("print once: id = %d\n", i)
		}
		go func() {
			once.Do(onceOperation)
			wg.Done()
		}()
	}

	wg.Wait()

}
