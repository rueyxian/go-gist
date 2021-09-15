package main

import (
	"bytes"
	"fmt"
	"log"
	"os/exec"
	"strings"
	"time"
)

func main() {

	// cmd := exec.Command("curl", "https://dl.google.com/go/go1.15.6.linux-amd64.tar.gz")
	cmd := exec.Command("go", "run", "./prog")
	var stdoutProcessStatus bytes.Buffer
	var stderr bytes.Buffer
	// cmd.Stdout = io.MultiWriter(ioutil.Discard, &stdoutProcessStatus)
	cmd.Stdout = &stdoutProcessStatus
	cmd.Stderr = &stderr

	fmt.Println("lol", stderr.String())

	done := make(chan struct{})
	go func() {
		tick := time.NewTicker(500 * time.Millisecond)
		defer tick.Stop()
		for {
			select {
			case <-done:
				return
			case <-tick.C:
				if strings.TrimSpace(stderr.String()) != "" {
					log.Print("error: %d", stderr.String())
				}

				log.Printf("downloaded: %d", stdoutProcessStatus.Len())
				// log.Printf("downloaded: %d", stdoutProcessStatus.String())
			}
		}
	}()
	err := cmd.Run()
	if err != nil {
		log.Fatalf("failed to call Run(): %v", err)
	}
	close(done)
}
