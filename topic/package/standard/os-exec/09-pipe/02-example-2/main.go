package main

import (
	"os"
	"os/exec"
)

func main() {

	cmd1 := exec.Command("cat", "fruits.txt")
	stdout1, err := cmd1.StdoutPipe()
	if err != nil {
		panic(err)
	}

	cmd2 := exec.Command("wc", "-l")
	cmd2.Stdin = stdout1
	cmd2.Stdout = os.Stdout
	cmd2.Stderr = os.Stderr

	if err := cmd1.Start(); err != nil {
		panic(err)
	}

	if err := cmd2.Start(); err != nil {
		panic(err)
	}

	cmd1.Wait()
	cmd2.Wait()

}
