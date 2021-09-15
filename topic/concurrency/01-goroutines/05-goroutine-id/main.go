package main

import (
	"bytes"
	"fmt"
	"runtime"
	"strconv"
	"time"
)

// source:
// https://blog.sgmansfield.com/2015/12/goroutine-ids/

func main() {
	fmt.Println(goid())

	go justAFunc()
	go justAFunc()
	go justAFunc()

	time.Sleep(1 * time.Second)

}

func justAFunc() {
	fmt.Println(goid())
}

func goid() uint64 {
	b := make([]byte, 64)
	b = b[:runtime.Stack(b, false)]
	b = bytes.TrimPrefix(b, []byte("goroutine "))
	b = b[:bytes.IndexByte(b, ' ')]
	n, _ := strconv.ParseUint(string(b), 10, 64)
	return n
}
