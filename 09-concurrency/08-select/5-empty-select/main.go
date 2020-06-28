package main

import "fmt"

func mario() {
	fmt.Println("hello, it's-a me, mario")
}

func luigi() {
	fmt.Println("hello, it's-a me, luigi")
}

func main() {
	fmt.Println("main: start")

	go mario()
	go luigi()

	// other goroutines will get executed, but after that main goroutine will stucks inside the select statement - deadlock.
	select {}

	fmt.Println("main: end")

}
