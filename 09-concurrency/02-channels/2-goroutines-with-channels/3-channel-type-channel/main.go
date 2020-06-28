package main

import "fmt"

func pika(ch chan string) {
	ch <- "pika pika~"
}

func pikachu(chch chan chan string) {
	ch := make(chan string)
	chch <- ch
}

func main() {

	chch := make(chan chan string)

	go pikachu(chch)

	ch := <-chch

	go pika(ch)

	fmt.Println(<-ch)

}
