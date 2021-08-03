package main

import "fmt"

func main() {

	sendCh := sendNums(1, 3, 5, 7, 9)
	recvCh := recvNums()

	for n := range sendCh {
		recvCh <- n
	}
	close(recvCh)

	fmt.Println("====================")
	// ====================
}

func sendNums(nums ...int) <-chan int {
	ret := make(chan int, len(nums))

	for _, n := range nums {
		fmt.Println("send", n)
		ret <- n
	}
	close(ret)

	return ret
}

func recvNums() chan<- int {
	ret := make(chan int)
	go func() {
		for n := range ret {
			fmt.Println("recv", n)
		}
	}()

	// for n := range ret {
	//   fmt.Println("recv", n)
	// }
	// close(ret)

	return ret
}
