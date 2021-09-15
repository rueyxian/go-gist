package main

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

// ================================================================================

type Data struct {
	Field string `validate:"required"`
	A     string
	B     string
}

// ================================================================================
var validate *validator.Validate

func main() {

	validate = validator.New()

	// ==================================================
	{
		data := Data{
			Field: "not zero",
			A:     "whatever",
			B:     "whatever",
		}

		if err := validate.Struct(data); err != nil {
			fmt.Println("=== 1 ===")
			printErr(err)
		}
	}

	// ====================
	{
		data := Data{
			Field: "",
			A:     "whatever",
			B:     "whatever",
		}

		if err := validate.Struct(data); err != nil {
			fmt.Println("=== 2 ===")
			printErr(err)
		}
	}

	// ==================================================

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
