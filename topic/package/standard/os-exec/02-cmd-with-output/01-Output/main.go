package main

import (
	"fmt"
	"os/exec"
)

func main() {

	cmd := exec.Command("go", "run", "./indr")

	// cmd.Output() does not capture stderr
	out, err := cmd.Output()
	if err != nil {
		panic(err)
	}

	fmt.Println(string(outerr))

}
