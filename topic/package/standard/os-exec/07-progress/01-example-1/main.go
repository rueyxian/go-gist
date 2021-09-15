package main

import (
	"bytes"
	"io"
	"io/ioutil"
	"log"
	"os/exec"
	"time"
)

func main() {

	// cmd := exec.Command("curl", "https://dl.google.com/go/go1.15.6.linux-amd64.tar.gz")
	cmd := exec.Command("curl", "https://dl.google.com/go/go1.15.6.linux-amd64.tar.gz")
	var stdoutProcessStatus bytes.Buffer
	cmd.Stdout = io.MultiWriter(ioutil.Discard, &stdoutProcessStatus)
	done := make(chan struct{})
	go func() {
		tick := time.NewTicker(time.Second)
		defer tick.Stop()
		for {
			select {
			case <-done:
				return
			case <-tick.C:
				log.Printf("downloaded: %d", stdoutProcessStatus.Len())
			}
		}
	}()
	err := cmd.Run()
	if err != nil {
		log.Fatalf("failed to call Run(): %v", err)
	}
	close(done)
}
