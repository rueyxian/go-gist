package main

import (
	"bytes"
	"fmt"
	"os/exec"
)

func main() {

	{
		cmd := exec.Command("go", "run", "./indr")

		var stdout, stderr bytes.Buffer
		cmd.Stdout = &stdout
		cmd.Stderr = &stderr

		if err := cmd.Run(); err != nil {
			panic(err)
		}

		fmt.Printf("stdout: \n%s\n", stdout.String())
		fmt.Printf("stderr: \n%s\n", stderr.String())
		fmt.Println("============================================================")
	}

	{
		cmd := exec.Command("go", "run", "./indr")

		// Or we can simply assign cmd.Stdout & cmd.Stderr
		// to literal &bytes.Buffer{}, this works too
		cmd.Stdout = &bytes.Buffer{}
		cmd.Stderr = &bytes.Buffer{}

		if err := cmd.Run(); err != nil {
			panic(err)
		}

		fmt.Printf("stdout: \n%s\n", cmd.Stdout)
		fmt.Printf("stderr: \n%s\n", cmd.Stderr)
		fmt.Println("============================================================")
	}

}
