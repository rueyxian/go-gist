package main

import (
	"errors"
	"runtime"
)

func main() {

	if err := isWindow(); err != nil {
		panic(err)
	}

}

func isWindow() error {
	if runtime.GOOS == "windows" {
		return errors.New("can't run on windows")
	}
	return nil
}
