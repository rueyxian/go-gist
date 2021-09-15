package main

import (
	"log"
	"os/exec"
)

func main() {

	// We can check the command that you wanted to execute
	// exists or not.
	path, err := exec.LookPath("go")
	if err != nil {
		log.Printf("'go' not found")
	} else {
		log.Printf("'go' is in '%s'\n", path)
	}

}
