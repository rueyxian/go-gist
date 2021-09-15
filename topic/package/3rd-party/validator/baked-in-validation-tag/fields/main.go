package main

import (
	"fmt"
	"time"

	"github.com/go-playground/validator"
)

var validate *validator.Validate

func main() {

	validate = validator.New()

	{
		a := time.Duration(1500 * time.Millisecond)
		b := time.Duration(1501 * time.Millisecond)
		if err := validate.VarWithValue(a, b, "eqfield"); err != nil {
			printErr(err)
		}
	}

}

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
