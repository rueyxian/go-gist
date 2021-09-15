package main

import (
	"fmt"
	"time"
)

func main() {

	done := make(chan struct{})

	go func() {

		fmt.Println("pause for 3 second")
		time.Sleep(3 * time.Second)
		close(done)

	}()

	// sending channel cannot unblocked
	// because you can't send value on a closed channel
	// error occurs
	done <- struct{}{}

	fmt.Println("done")

}
