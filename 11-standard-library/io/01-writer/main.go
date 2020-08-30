package main

import (
	"fmt"
)

type myWriter struct {
	data []byte
}

func (w *myWriter) Write(p []byte) (int, error) {
	n := len(p)
	w.data = make([]byte, n)
	copy(w.data, p)
	return n, nil
}

func main() {

	w := new(myWriter)

	n, err := fmt.Fprintln(w, "something stupid")

	if err != nil {
		panic(err)
	}

	fmt.Println(n, "bytes written.")
	fmt.Println(w.data)

}
