package main

/*
If a goroutine is sending data on a channel,
some other goroutine is expected to receive that data.
Else dealock occurs
*/

func main() {
	ch := make(chan int)
	ch <- 123
}
