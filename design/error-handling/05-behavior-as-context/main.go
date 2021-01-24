package main

import (
	"fmt"
	"sandbox/go-jottings/design/error-handling/dummy"
)

type client struct {
}

func (c client) BehaviorAsContext() error {
	err := dummy.Call(c)

	if err != nil {
		switch e := err.(type) {

		case dummy.Fatal:
			if e.Fatal() {
				return err
			}
		// case *dummy.AlphaError:
		//   if e.Fatal() {
		//     return err
		//   }

		// case *dummy.BetaError:
		//   if e.Fatal() {
		//     return err
		//   }

		// case *dummy.GammaError:
		//   if e.Fatal() {
		//     return err
		//   }
		default:
		}

	}
	return nil
}

func main() {

	var c client

	if e := c.BehaviorAsContext(); e != nil {
		fmt.Println(e)
		return
	}

	fmt.Println("peace out =)")
}
