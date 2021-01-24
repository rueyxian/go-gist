package main

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

var (
	ErrBadRequest = errors.New("bad request")
	ErrTimeout    = errors.New("timeout")
)

func main() {

	if err := webCall(); err != nil {

		switch err {
		case ErrBadRequest:
			fmt.Println("Bad Request Occurred")
			return
		case ErrTimeout:
			fmt.Println("Request Timeout")
			return
		default:
			fmt.Println(err)
			return
		}
	}

	fmt.Println("^.^")
}

func webCall() error {
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	var err error
	switch rnd.Intn(3) {
	case 1:
		err = ErrBadRequest
	case 2:
		err = ErrTimeout
	default:
	}

	return err
}
