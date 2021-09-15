package main

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

type User struct {
	FirstName string `validate:"required_without=LastName"`
	LastName  string `validate:"required_without=FirstName"`
	Age       int    `validate:""`
}

// ================================================================================

var validate *validator.Validate

func main() {

	validate = validator.New()

	{
		user := &User{
			FirstName: "Lorem",
		}
		if err := validate.Struct(user); err != nil {
			fmt.Println("=== 1 ===")
			printErr(err)
		}
	}
	// ==============================
	{
		user := &User{
			LastName: "Ipsum",
		}
		if err := validate.Struct(user); err != nil {
			fmt.Println("=== 2 ===")
			printErr(err)
		}
	}
	// ==============================
	{
		user := &User{
			Age: 100,
		}
		if err := validate.Struct(user); err != nil {
			fmt.Println("=== 3 ===")
			printErr(err)
		}
	}

}

// ================================================================================
// printErr
func printErr(err error) {
	for _, err := range err.(validator.ValidationErrors) {
		fmt.Println("Namespace      :", err.Namespace())
		fmt.Println("Field          :", err.Field())
		fmt.Println("StructNamespace:", err.StructNamespace())
		fmt.Println("StructField    :", err.StructField())
		fmt.Println("Tag            :", err.Tag())
		fmt.Println("ActualTag      :", err.ActualTag())
		fmt.Println("Kind           :", err.Kind())
		fmt.Println("Type           :", err.Type())
		fmt.Println("Value          :", err.Value())
		fmt.Println("Param          :", err.Param())
		fmt.Println()
	}
}
