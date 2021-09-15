package main

import (
	"fmt"
	"log"
	"os/exec"
)

// https://unix.stackexchange.com/questions/99263/what-does-21-in-this-command-mean

// https://unix.stackexchange.com/questions/746/how-could-i-remember-how-to-use-redirection

// unix commands note:
// semi-colon ";" is command seperator

func main() {
	cmd := exec.Command("sh", "-c", "echo stdout; echo 1>&2 stderr")
	stdoutStderr, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", stdoutStderr)
}
