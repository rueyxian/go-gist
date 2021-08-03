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

	// receiving channel will be unblocked once it is closed
	<-done

	fmt.Println("done")

}
