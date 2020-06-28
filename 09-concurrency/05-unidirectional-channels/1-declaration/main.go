package main

import "fmt"

func main() {

	ch := make(chan int)     // bidirectional channel
	chSO := make(chan<- int) //send-only channel
	chRO := make(<-chan int) //receive-only channel

	fmt.Printf("bidirectional: \t %T \n", ch)
	fmt.Printf("send-only: \t %T \n", chSO)
	fmt.Printf("receive-only: \t %T \n", chRO)

	fmt.Println("")


	
}
