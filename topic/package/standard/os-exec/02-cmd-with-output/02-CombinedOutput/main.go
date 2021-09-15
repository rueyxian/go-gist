package main

import (
	"fmt"
	"os/exec"
)

func main() {

	cmd := exec.Command("go", "run", "./indr")

	// cmd.CombinedOutput() capture both stdout and stderr
	outerr, err := cmd.CombinedOutput()
	if err != nil {
		panic(err)
	}

	fmt.Println(string(outerr))

}

// CombinedOutput method

// func (c *Cmd) CombinedOutput() ([]byte, error) {
//   if c.Stdout != nil {
//     return nil, errors.New("exec: Stdout already set")
//   }
//   if c.Stderr != nil {
//     return nil, errors.New("exec: Stderr already set")
//   }
//   var b bytes.Buffer
//   c.Stdout = &b
//   c.Stderr = &b
//   err := c.Run()
//   return b.Bytes(), err
// }
