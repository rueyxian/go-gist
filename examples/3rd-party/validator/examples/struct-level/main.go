package main

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/go-playground/validator/v10"
)

// ================================================================================
// User contains user information
type User struct {
	FirstName string `json:"fname" validate:"required"`
	LastName  string `json:"lname" validate:"required"`
	// FirstName      string     `json:"fname"`
	// LastName       string     `json:"lname"`
	Age            uint8      `validate:"gte=0,lte=130"`
	Email          string     `json:"e-mail" validate:"required,email"`
	FavouriteColor string     `validate:"hexcolor|rgb|rgba"`
	Addresses      []*Address `validate:"required,dive,required"` // a person can have a home and cottage...
}

// Address houses a users address information
type Address struct {
	Street string `validate:"required"`
	City   string `validate:"required"`
	Planet string `validate:"required"`
	Phone  string `validate:"required"`
}

// ================================================================================
var validate *validator.Validate

func main() {

	validate = validator.New()

	// ==============================

	validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
		if name == "-" {
			return ""
		}
		return name
	})

	// validate.RegisterStructValidation(UserStructLevelValidation, User{})

	// ==============================

	// build 'User' info, normally posted data etc...
	address := &Address{
		Street: "123 Sesame Street",
		Planet: "Aiur",
		Phone:  "none",
		City:   "Unknown",
	}

	user := &User{
		FirstName:      "",
		LastName:       "",
		Age:            45,
		Email:          "Badger.Smith@gmail",
		FavouriteColor: "#000",
		Addresses:      []*Address{address},
	}

	// ==============================

	if err := validate.Struct(user); err != nil {
		printErr(err)
	}

	// ==============================

	// rvf := []reflect.StructField{
	//   {
	//     Name: "Name",
	//     Type: reflect.TypeOf(""),
	//     Tag:  `json:"name,namee"`,
	//   },
	//   {
	//     Name: "Price",
	//     Type: reflect.TypeOf(float64(0)),
	//     Tag:  `json:"px"`,
	//   },
	// }

	// for _, f := range rvf {
	//   JsonTagName(f)
	// }

}

// ================================================================================
func JsonTagName(fld reflect.StructField) string {

	// fmt.Println(fld.Tag.Get("json"))

	name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
	fmt.Println(name)

	return ""
	// ==============================

	// name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
	// if name == "-" {
	//   return ""
	// }
	// return name
}

// ================================================================================

func UserStructLevelValidation(sl validator.StructLevel) {

	user := sl.Current().Interface().(User)

	if len(user.FirstName) == 0 && len(user.LastName) == 0 {
		sl.ReportError(user.FirstName, "fname", "FirstName", "fnameorlname", "")
		sl.ReportError(user.LastName, "lname", "LastName", "fnameorlname", "")
	}

	// plus can do more, even with different tag than "fnameorlname"
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
