package main

import (
	"fmt"

	"github.com/go-playground/validator"
)

// ================================================================================
// User
type User struct {
	FirstName      string     `validate:"required"`
	LastName       string     `validate:"required"`
	Age            uint8      `validate:"gte=0,lte=130"`
	Email          string     `validate:"required,email"`
	FavouriteColor string     `validate:"iscolor"`
	Addresses      []*Address `validate:"required,dive,required"`
}

// Address
type Address struct {
	Street string `validate:"required"`
	City   string `validate:"required"`
	Planet string `validate:"required"`
	Phone  string `validate:"required"`
}

// ================================================================================
var validate *validator.Validate

// main
func main() {
	validate = validator.New()

	validateStruct()
	validateVariable()
}

// ================================================================================
func validateStruct() {
	address1 := &Address{
		Street: "123 Sesame Street",
		Planet: "Aiur",
		Phone:  "none",
	}

	address2 := &Address{
		Street: "E18B Cybertown",
		City:   "Poop",
		Planet: "Moon",
	}

	user := &User{
		FirstName:      "Lorem",
		LastName:       "Ipsum",
		Age:            135,
		Email:          "Lorem.Ipsum@gmail.com",
		FavouriteColor: "#000-",
		Addresses:      []*Address{address1, address2},
	}

	err := validate.Struct(user)
	if err != nil {
		printErr(err)
	}
}

// ================================================================================
func validateVariable() {
	v1 := "jackjazzrabit.gmail.com"
	if err := validate.Var(v1, "required,email"); err != nil {
		printErr(err)
	}

	v2 := 0 // <- ensure not zero value
	if err := validate.Var(v2, "required"); err != nil {
		printErr(err)
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
