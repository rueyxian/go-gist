package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	const name, age = "Kim", 22
	// n, err := fmt.Fprintf(os.Stdout, "%s is %d years old.\n", name, age)

	// // The n and err return values from Fprintf are
	// // those returned by the underlying io.Writer.
	// if err != nil {
	// 	fmt.Fprintf(os.Stderr, "Fprintf: %v\n", err)
	// }
	// fmt.Printf("%d bytes written.\n", n)

	file, _ := os.Create("file.txt")
	w := bufio.NewWriter(file)

	n, err := fmt.Fprintf(w, "%s is %d years old.\n", name, age)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fprintf: %v\n", err)
	}
	fmt.Printf("%d bytes written.\n", n)

	w.Flush()

}
