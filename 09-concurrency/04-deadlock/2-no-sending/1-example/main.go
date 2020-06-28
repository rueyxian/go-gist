package main

import "fmt"

/*
	By the same token, if a goroutine is waiting for receiving from a channel, some other goroutine is expected to send data, else deadlock occurs
*/

func main() {
	ch := make(chan int)

	fmt.Println(<-ch)
}
