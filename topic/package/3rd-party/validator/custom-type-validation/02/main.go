package main

import (
	"database/sql"
	"database/sql/driver"
	"fmt"
	"reflect"

	"github.com/go-playground/validator/v10"
)

type User struct {
	Name sql.NullString `validate:"required"`
	Age  sql.NullInt64  `validate:"required"`
}

// ================================================================================

var validate *validator.Validate

func main() {

	validate = validator.New()

	validate.RegisterCustomTypeFunc(UserCustomType,
		sql.NullString{},
		sql.NullInt64{},
		sql.NullBool{},
		sql.NullFloat64{},
	)

	// ==============================
	{
		user := &User{
			Name: sql.NullString{String: "Lorem", Valid: true},
			Age:  sql.NullInt64{Int64: 120, Valid: true},
		}

		if err := validate.Struct(user); err != nil {
			fmt.Println("=== 1 ===")
			printErr(err)
		}
	}
	// ==============================
	{
		user := &User{
			Name: sql.NullString{String: "", Valid: true},
			Age:  sql.NullInt64{Int64: 0, Valid: true},
		}

		if err := validate.Struct(user); err != nil {
			fmt.Println("=== 2 ===")
			printErr(err)
		}
	}
	// ==============================
	{
		user := &User{
			Name: sql.NullString{String: "", Valid: true},
			Age:  sql.NullInt64{Int64: 17, Valid: true},
		}

		if err := validate.Struct(user); err != nil {
			fmt.Println("=== 3 ===")
			printErr(err)
		}
	}

}

// ================================================================================
func UserCustomType(field reflect.Value) interface{} {

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
