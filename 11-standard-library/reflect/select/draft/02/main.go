package main

import "fmt"

func main() {



	// ch := make(chan int)

	// go func(c chan int) {
	//   for i := 0; i < 10; i++ {
	//     ch <- i
	//   }
	//   close(c)
	// }(ch)

	// selCase := []reflect.SelectCase{
	//   {Dir: reflect.SelectRecv, Chan: reflect.ValueOf(ch)},
	// }

	// for i := 0; i < 1; {
	//   idx, v, ok := reflect.Select(selCase)
	//   if ok {
	//     fmt.Println(idx, v.Int(), ok)
	//   } else {
	//     i++
	//   }

	// }

	s := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	fmt.Println(s[:0])
	fmt.Println(s[4:])

}

func
