package main

import "fmt"

func main() {

	ns := []int{1, 2, 4, 5, 7}
	operation(ns...)

}

func operation(nums ...int) {
	for num := range nums {
		fmt.Println(num)
	}
}
