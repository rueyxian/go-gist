package main

import (
	"errors"
	"fmt"
	"os/exec"
	"runtime"
)

func main() {
	if err := isWindow(); err != nil {
		panic(err)
	}
	execute()
}

func execute() {
	out, err := exec.Command("ls", "-ltr").Output()
	if err != nil {
		panic(err)
	}
	fmt.Println(string(out))
}

func isWindow() error {
	if runtime.GOOS == "windows" {
		return errors.New("can't run on windows")
	}
	return nil
}
