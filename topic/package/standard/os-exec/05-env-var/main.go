package main

import (
	"fmt"
	"os/exec"
)

// ==============================

func main() {
	cmd := exec.Command("bash", "-c", "echo $myvar")
	cmd.Env = []string{"myvar=abc"}
	outerr, err := cmd.CombinedOutput()
	if err != nil {
		panic(err)
	}
	fmt.Println(string(outerr))

}
