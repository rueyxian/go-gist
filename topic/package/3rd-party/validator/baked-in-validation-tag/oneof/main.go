package main

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

func main() {

	validate = validator.New()

	{
		v := "corgi"
		if err := validate.Var(v, `oneof='corgi' 'samoyad' 'golden retriever'`); err != nil {
			printErr(err)
		}
	}

	{
		v := "papillion"
		if err := validate.Var(v, `oneof='corgi' 'samoyad' 'golden retriever'`); err != nil {
			printErr(err)
		}
	}

}

// ================================================================================

// func printErr(err error) {
//   noempty := func(name string, a interface{}) {
//     switch v := a.(type) {
//     case func() string, func() reflect.Kind, func() reflect.Type, func() interface{}:
//       rvs := reflect.ValueOf(v).Call([]reflect.Value{})
//       if s := rvs[0].Interface(); s != "" {
//         fmt.Println(name, s)
//       }
//     default:
//       panic("printErr: invalid type")
//     }
//   }

//   for _, err := range err.(validator.ValidationErrors) {
//     noempty("Namespace      :", err.Namespace)
//     noempty("Field          :", err.Field)
//     noempty("StructNamespace:", err.StructNamespace)
//     noempty("StructField    :", err.StructField)
//     noempty("Tag            :", err.Tag)
//     noempty("ActualTag      :", err.ActualTag)
//     noempty("Kind           :", err.Kind)
//     noempty("Type           :", err.Type)
//     noempty("Value          :", err.Value)
//     noempty("Param          :", err.Param)
//     fmt.Println()
//   }
// }

// ================================================================================

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
	}
}
