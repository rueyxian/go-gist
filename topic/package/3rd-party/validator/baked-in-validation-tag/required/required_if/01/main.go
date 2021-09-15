package main

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

// ================================================================================

type Data struct {
	Field string `validate:"required_if=A foo"`
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
			A:     "foo",
		}

		if err := validate.Struct(data); err != nil {
			fmt.Println("=== 1 ===")
			printErr(err)
		}
	}

	// ====================
	{
		data := Data{
			Field: "not zero",
			A:     "whatever",
		}

		if err := validate.Struct(data); err != nil {
			fmt.Println("=== 2 ===")
			printErr(err)
		}
	}

	// ==================================================

	{
		data := Data{
			Field: "",
			A:     "not foo",
		}

		if err := validate.Struct(data); err != nil {
			fmt.Println("=== 3 ===")
			printErr(err)
		}
	}

	// ====================

	{
		data := Data{
			Field: "whatever",
			A:     "not foo",
		}

		if err := validate.Struct(data); err != nil {
			fmt.Println("=== 4 ===")
			printErr(err)
		}
	}

	// ==================================================

	{
		data := Data{
			Field: "",
			A:     "foo",
		}

		if err := validate.Struct(data); err != nil {
			fmt.Println("=== 5 ===")
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
