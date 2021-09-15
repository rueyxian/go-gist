package main

import (
	"fmt"
)

func main() {

	// time.Sleep(time.Millisecond * 500)
	// fmt.Println("num of goroutines: ", runtime.NumGoroutine())

	sendCh := sendNums(1, 3, 5, 7, 9)

	// time.Sleep(time.Millisecond * 500)
	// fmt.Println("num of goroutines: ", runtime.NumGoroutine())

	recvCh := recvNums()

	// time.Sleep(time.Millisecond * 500)
	// fmt.Println("num of goroutines: ", runtime.NumGoroutine())

	for n := range sendCh {
		recvCh <- n
	}
	close(recvCh)

	// time.Sleep(time.Millisecond * 500)
	// fmt.Println("num of goroutines: ", runtime.NumGoroutine())

	fmt.Println("====================")
	// ====================
}

func sendNums(nums ...int) <-chan int {
	ret := make(chan int)
	go func() {
		for _, n := range nums {
			fmt.Println("send", n)
			ret <- n
		}
		close(ret)
	}()
	return ret
}

func recvNums() chan<- int {
	ret := make(chan int)
	go func() {
		for n := range ret {
			fmt.Println("recv", n)
		}
	}()
	return ret
}
