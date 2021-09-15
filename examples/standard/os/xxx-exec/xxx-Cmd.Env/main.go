package main

import (
	"errors"
	"fmt"
	"os"
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
	{
		out, err := exec.Command("go", "run", "./foo/main.go").Output()
		if err != nil {
			panic(err)
		}
		fmt.Println(string(out))
	}

	{
		//TODO: diff between Run() and Output()

		// cmd := exec.Command("go", "run", "./foo/main.go")
		cmd := exec.Command("go", "run", "./foo/main.go")
		cmd.Env = append(os.Environ(), "name=solar", "age=27")

		// if err := cmd.Run(); err != nil {
		//   panic(err)
		// }

		out, err := cmd.CombinedOutput()
		if err != nil {
			panic(err)
		}

		fmt.Println(string(out))

	}
}

func isWindow() error {
	if runtime.GOOS == "windows" {
		return errors.New("can't run on windows")
	}
	return nil
}
