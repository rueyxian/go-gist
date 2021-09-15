package main

import (
	"os"
	"os/exec"
)

func main() {
	stdin, err := os.Open("main.go")
	if err != nil {
		panic(err)
	}
	cmd := exec.Command("wc", "-l")
	cmd.Stdin = stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err = cmd.Run()
	if err != nil {
		panic(err)
	}
}
