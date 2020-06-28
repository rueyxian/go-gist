package main

import (
	"fmt"
)

/*
number = 1234
squares = (1*1)+(2*2)+(3*3)+(4*4*4)
cubes = (1*1*1)+(2*2*2)+(3*3*3)+(4*4*4)
*/

func calSquares(n int, ch chan int) {
	sum := 0
	for n != 0 {
		d := n % 10
		sum += d * d
		n /= 10
	}
	ch <- sum
}

func calCubes(n int, ch chan int) {
	sum := 0
	for n != 0 {
		d := n % 10
		sum += d * d * d
		n /= 10
	}
	ch <- sum
}

func main() {
	number := 4896
	chSqr := make(chan int)
	chCbe := make(chan int)
	go calSquares(number, chSqr)
	go calCubes(number, chCbe)
	sqr, cbe := <-chSqr, <-chCbe
	fmt.Printf("squares: \t %v \n", sqr)
	fmt.Printf("  cubes: \t %v \n", cbe)
}
