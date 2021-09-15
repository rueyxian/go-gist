package main

import (
	"fmt"
	"os/exec"
)

func main() {

	cmd1 := exec.Command("cat", "fruits.txt")
	cmd2 := exec.Command("wc", "-l")

	data, err := pipeCommands(cmd1, cmd2)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(data))

}

// ================================================================================

func pipeCommands(cmds ...*exec.Cmd) ([]byte, error) {
	for i, cmd := range cmds[:len(cmds)-1] {
		stdout, err := cmd.StdoutPipe()
		if err != nil {
			return nil, err
		}
		cmds[i+1].Stdin = stdout
		if err := cmd.Start(); err != nil {
			return nil, err
		}
	}
	cmdN, err := cmds[len(cmds)-1].CombinedOutput()
	if err != nil {
		return nil, err
	}
	return cmdN, nil
}
