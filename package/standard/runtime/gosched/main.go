package main

import (
	"fmt"
	"runtime"
	"strconv"
	"strings"
)

func init() {
	// runtime.GOMAXPROCS(2)
}

func goid() int {
	var buf [64]byte
	n := runtime.Stack(buf[:], false)
	idField := strings.Fields(strings.TrimPrefix(string(buf[:n]), "goroutine "))[0]
	id, err := strconv.Atoi(idField)
	if err != nil {
		panic(fmt.Sprintf("cannot get goroutine id: %v", err))
	}
	return id
}

func say(s string) {
	for i := 0; i < 2; i++ {
		runtime.Gosched()
		fmt.Printf("goid: %2d \t %s\n", goid(), s)

		// wg.Done()
		// time.Sleep(time.Millisecond * 10)
	}
}

func main() {

	// n := 5

	// var wg sync.WaitGroup
	// wg.Add(n)

	go say("sub")
	say("main")

	// wg.Wait()
}
