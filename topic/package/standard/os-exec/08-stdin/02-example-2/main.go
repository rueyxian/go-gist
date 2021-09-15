package main

import (
	"os"
	"os/exec"
)

func main() {

	cmd := exec.Command("wc", "file.txt")

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		panic(err)
	}

}
