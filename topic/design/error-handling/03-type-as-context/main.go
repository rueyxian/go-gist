package main

import (
	"fmt"
	"go-gist/design/error-handling/04-type-as-context/config"
)

type cfg struct {
	User     string
	Password string
	Secure   bool
}

func main() {

	var c cfg
	_ = c
	// s := []string{"User=nova", "Password=4896", "Secure=true"}
	s := []string{"User=nova", "Drowssap=4896", "Secure=true"}

	// ==============================
	// if err := config.Parse(&c, s); err != nil {
	// if err := config.Parse(c, s); err != nil {
	// if err := config.Parse((*cfg)(nil), s); err != nil {
	if err := config.Parse((*cfg)(nil), s); err != nil {
		switch err.(type) {
		case *config.InvalidParseError:
			fmt.Printf("InvalidParseError: %s\n", err)
		case *config.InvalidArgumentError:
			fmt.Printf("InvalidArgumentError: %s\n", err)
		default:
			fmt.Println(err)
		}
		return
	}

}
