package main

import (
	"fmt"
	"os/exec"
	"runtime"
)

func main() {

	if isWindow() {
		fmt.Println("can't run on windows")
		return
	}

	cmd := exec.Command("ls", "-ltra")
	err := cmd.Run()
	if err != nil {
		panic(err)
	}

	// If we run the program, it won't print anything,
	// though, it does run something already.

	// cmd.Run() is a wrapper of cmd.Start() & cmd.Wait() method:
	// func (c *Cmd) Run() error {
	//   if err := c.Start(); err != nil {
	//     return err
	//   }
	//   return c.Wait()
	// }

	// So if you want more fine-grain control, use cmd.Start() & cmd.Wait()

}

func isWindow() bool {
	return runtime.GOOS == "windows"
}
