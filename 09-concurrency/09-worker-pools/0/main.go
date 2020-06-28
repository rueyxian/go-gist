package main

import (
	"fmt"
	"sync"
	"time"
)

func process(i int, wg *sync.WaitGroup) {
	fmt.Printf("start  %d \n", i)
	time.Sleep(2 * time.Second)
	fmt.Printf("end    %d \n", i)
	wg.Done()
}

func main() {
	// no := 5
	// var wg sync.WaitGroup	// empty struct wait group

	// fmt.Printf("%T \t %+v \n", wg, wg)

	// for i := 0; i < no; i++ {
	// 	wg.Add(1)
	// 	go process(i, &wg)
	// }
	// wg.Wait()
	// fmt.Println("All go routines finished executing")

	// ============================================================

	fmt.Printf("main: start \n")
	var wg sync.WaitGroup
	// wg.Add(3)
	time.Sleep(1000 * time.Millisecond)
	// wg.Add(-1)
	// wg.Add(-1)
	wg.Done()

	wg.Wait()
	fmt.Printf("main:   end \n")

}

func f(wg *sync.WaitGroup) {
	time.Sleep(4000 * time.Millisecond)
	wg.Done()
}
