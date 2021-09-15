package main

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

type NullString struct {
	String string
	Valid  bool
}

type NullInt64 struct {
	Int64 int64
	Valid bool
}

type NullBool struct {
	Bool  bool
	Valid bool
}

// ================================================================================

type User struct {
	Name       NullString
	Age        NullInt64
	Vaccinated NullBool
}

// ================================================================================

var validate *validator.Validate

func main() {

	validate = validator.New()

	validate.RegisterStructValidation(UserStructLevelValidation, User{})

	// ==============================
	{
		user := &User{
			Name:       NullString{String: "Lorem", Valid: true},
			Age:        NullInt64{Int64: 120, Valid: true},
			Vaccinated: NullBool{Bool: false, Valid: true},
		}
		if err := validate.Struct(user); err != nil {
			fmt.Println("=== 1 ===")
			printErr(err)
		}
	}
	// ==============================
	{
		user := &User{
			Name:       NullString{String: "Ipsum", Valid: false},
			Age:        NullInt64{Int64: 0, Valid: true},
			Vaccinated: NullBool{Bool: true, Valid: true},
		}
		if err := validate.Struct(user); err != nil {
			fmt.Println("=== 2 ===")
			printErr(err)
		}
	}
	// ==============================
	{
		user := &User{
			Name:       NullString{String: "Dolor", Valid: true},
			Age:        NullInt64{Int64: 256, Valid: false},
			Vaccinated: NullBool{Bool: false, Valid: true},
		}
		if err := validate.Struct(user); err != nil {
			fmt.Println("=== 3 ===")
			printErr(err)
		}
	}

}

// ================================================================================
func UserStructLevelValidation(sl validator.StructLevel) {
	user := sl.Current().Interface().(User)

	if !user.Name.Valid {
		sl.ReportError(user.Name.Valid, "Name", "user.Name", "is_valid", "")
	}

	if !user.Age.Valid {
		sl.ReportError(user.Age.Valid, "Age", "user.Age", "is_valid", "")
	}

	if !user.Vaccinated.Valid {
		sl.ReportError(user.Vaccinated.Valid, "Vaccinated", "user.Vaccinated", "is_valid", "")
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
