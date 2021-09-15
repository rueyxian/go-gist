package main

import (
	"fmt"

	"github.com/go-playground/locales/en"
	"github.com/go-playground/validator/v10"

	ut "github.com/go-playground/universal-translator"
	en_translations "github.com/go-playground/validator/v10/translations/en"
)

// ================================================================================

var (
	uni      *ut.UniversalTranslator
	validate *validator.Validate
)

func main() {

	en := en.New()
	uni = ut.New(en, en)
	trans, _ := uni.GetTranslator("en")

	validate = validator.New()
	en_translations.RegisterDefaultTranslations(validate, trans)

	// translateAll(trans)
	translateIndividual(trans)
	// translateOverride(trans) // yep you can specify your own in whatever locale you want!

	// ==============================

}

// ================================================================================

// func translateAll(trans ut.Translator) {

// }

// func translateIndividual(trans ut.Translator) {

// }

// func translateOverride(trans ut.Translator) {

// }

// ================================================================================

func translateAll(trans ut.Translator) {

	type User struct {
		Username string `validate:"required"`
		Tagline  string `validate:"required,lt=10"`
		Tagline2 string `validate:"required,gt=1"`
	}

	user := User{
		Username: "Joeybloggs",
		Tagline:  "This tagline is way too long.",
		Tagline2: "1",
	}

	// if err := validate.Struct(user); err != nil {
	//   printErr(err)
	// }

	// ==============================

	if err := validate.Struct(user); err != nil {
		// translate all error at once
		errs := err.(validator.ValidationErrors)

		// returns a map with key = namespace & value = translated error
		// NOTICE: 2 errors are returned and you'll see something surprising
		// translations are i18n aware!!!!
		// eg. '10 characters' vs '1 character'
		fmt.Println(errs.Translate(trans))

	}

}

func translateIndividual(trans ut.Translator) {

	type User struct {
		Username string `validate:"required"`
	}

	var user User

	err := validate.Struct(user)
	if err != nil {

		errs := err.(validator.ValidationErrors)

		for _, e := range errs {
			// can translate each error one at a time.
			fmt.Println(e.Translate(trans))
		}
	}
}

func translateOverride(trans ut.Translator) {

	validate.RegisterTranslation("required", trans, func(ut ut.Translator) error {
		return ut.Add("required", "{0} must have a value!", true) // see universal-translator for details
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("required", fe.Field())

		return t
	})

	type User struct {
		Username string `validate:"required"`
	}

	var user User

	err := validate.Struct(user)
	if err != nil {

		errs := err.(validator.ValidationErrors)

		for _, e := range errs {
			// can translate each error one at a time.
			fmt.Println(e.Translate(trans))
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
