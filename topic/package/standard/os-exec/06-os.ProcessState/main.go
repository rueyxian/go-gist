package main

import (
	"log"
	"os"
	"os/exec"
)

func main() {

	cmd := exec.Command("go", "run", "./prog")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Start()
	if err != nil {
		panic(err)
	}

	log.Printf("pid: %d", cmd.Process.Pid)
	cmd.Process.Wait()
	log.Printf("exitcode: %d", cmd.ProcessState.ExitCode())

}

// cmd := exec.Command("bash", "-c", "sleep 1;echo $myvar")
// cmd.Stdout = os.Stdout
// cmd.Stderr = os.Stderr
// err := cmd.Start()
// if err != nil {
//   log.Fatalf("failed to call cmd.Start(): %v", err)
// }
// log.Printf("pid: %d", cmd.Process.Pid)
// cmd.Process.Wait()
// log.Printf("exitcode: %d", cmd.ProcessState.ExitCode())
