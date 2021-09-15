package main

import (
	"os"
	"os/exec"
)

func main() {
	cmdCat := exec.Command("cat", "main.go")
	catout, err := cmdCat.StdoutPipe()
	if err != nil {
		panic(err)
	}
	cmdWC := exec.Command("wc", "-l")
	cmdWC.Stdin = catout
	cmdWC.Stdout = os.Stdout
	err = cmdCat.Start()
	if err != nil {
		panic(err)
	}
	err = cmdWC.Start()
	if err != nil {
		panic(err)
	}
	cmdCat.Wait()
	cmdWC.Wait()
}
