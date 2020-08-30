package main

import "fmt"

func main() {
	var a chan int
	if a == nil {
		fmt.Println("channel a is nil, going to define it")
		a = make(chan int)
		fmt.Printf("Type of a is %T\n", a)
	}

	data1 := 4896
	a <- data1 // send to channel a

	data2 := <-a // receive from channel a
	<-a          // receive without assignment

	_ = data2

}
