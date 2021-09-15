package main

import (
	"database/sql"
	"database/sql/driver"
	"fmt"
	"reflect"

	"github.com/go-playground/validator"
)

// DbBackedUser User struct
type DbBackedUser struct {
	Name       sql.NullString `validate:"required"`
	Age        sql.NullInt64  `validate:"required"`
	Vaccinated sql.NullBool   `validate:"required"`
}

// ================================================================================

var validate *validator.Validate

func main() {

	validate = validator.New()

	// validate.RegisterCustomTypeFunc(ValidateValuer,
	//   sql.NullString{},
	//   sql.NullInt64{},
	//   sql.NullBool{},
	//   sql.NullFloat64{},
	// )

	v := DbBackedUser{
		Name:       sql.NullString{String: "nova", Valid: false},
		Age:        sql.NullInt64{Int64: 0, Valid: true},
		Vaccinated: sql.NullBool{Bool: false, Valid: true},
	}

	// if err := validate.Struct(v); err != nil {
	//   printErr(err)
	// }

	// ==============================

	fmt.Println(v.Name.Valid)
	val, err := v.Name.Value()
	if err != nil {
		panic(err)
	}
	fmt.Println(val)

}

// ================================================================================
// func ValidateValuer(field reflect.Value) interface{} {
//   return nil
// }

func ValidateValuer(field reflect.Value) interface{} {

	if valuer, ok := field.Interface().(driver.Valuer); ok {

		val, err := valuer.Value()
		if err == nil {
			return val
		}
		// handle the error how you want
	}

	return nil
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
