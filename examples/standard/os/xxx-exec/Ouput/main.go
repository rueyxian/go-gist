package main

import (
	"fmt"
	"os/exec"
)

func main() {

	out, err := exec.Command("date").Output()

	if err != nil {
		panic(err)
	}

	fmt.Println(string(out))

}
