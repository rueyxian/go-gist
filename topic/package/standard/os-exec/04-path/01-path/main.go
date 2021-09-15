package main

import (
	"log"
	"os/exec"
)

func main() {

	{
		// If the program $PATH is set,
		// path can be omitted when executing the program
		cmd := exec.Command("go", "version")
		log.Printf("path: %s", cmd.Path)
	}

	{
		// If the program $PATH is not set,
		// we have to provide the full path
		cmd := exec.Command("/usr/local/go/bin/go", "version")
		log.Printf("path: %s", cmd.Path)
	}

	{
		// Or we can set the cmd.Dir
		cmd := exec.Command("go", "version")
		cmd.Dir = "/usr/local/go/bin/go"
		log.Printf("path: %s", cmd.Path)
	}

}
