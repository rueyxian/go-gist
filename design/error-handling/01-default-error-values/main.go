package main

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

// http://golang.org/pkg/builtin/#error
// type error interface {
//   Error() string
// }

// ==========

// // http://golang.org/src/pkg/errors/errors.go
// type errorString struct {
//   s string
// }

// // http://golang.org/src/pkg/errors/errors.go
// func (e *errorString) Error() string {
//   return e.s
// }

// // http://golang.org/src/pkg/errors/errors.go
// func New(text string) error {
//   return &errorString{text}
// }

// ==========

func main() {

	if err := webCall(); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("^.^")
}

func webCall() error {
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	switch rnd.Intn(3) {
	case 1:
		return errors.New("bad request")
	case 2:
		return errors.New("timeout")
	default:
	}

	return nil
}
