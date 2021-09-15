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
	utranslate *ut.UniversalTranslator
	validate   *validator.Validate
)

// ================================================================================

type User struct {
	Name    string `validate:"required"`
	Level   int    `validate:"required,gt=0,lt=100"`
	Email   string `validate:"required,email"`
	Tagline string `validate:"required,lt=20"`
	Class   string `validate:"required"`
}

// ================================================================================

func main() {
	en := en.New()
	uni := ut.New(en)
	trans, _ := uni.GetTranslator("en")

	validate = validator.New()
	en_translations.RegisterDefaultTranslations(validate, trans)
	// en.RegisterDefaultTranslations(validate, trans)

	// ==============================
	{
		user := &User{
			Name:    "gopher_forever",
			Level:   120,
			Email:   "dumpsterdiver.gmail.com",
			Tagline: "The quick brown fox jumps over the lazy dog",
		}
		if err := validate.Struct(user); err != nil {
			fmt.Println("=== no translator ===")
			noTranslator(err)
		}
	}

	// ==============================
	fmt.Println("============================================================\n")
	// ==============================
	{
		user := &User{
			Name:    "gopher_forever",
			Level:   120,
			Email:   "dumpsterdiver.gmail.com",
			Tagline: "The quick brown fox jumps over the lazy dog",
		}
		if err := validate.Struct(user); err != nil {
			fmt.Println("=== translate all ===")
			translateAll(trans, err)
		}
	}

	// ==============================
	fmt.Println("============================================================\n")
	// ==============================
	{
		user := &User{
			Name:    "gopher_forever",
			Level:   120,
			Email:   "dumpsterdiver.gmail.com",
			Tagline: "The quick brown fox jumps over the lazy dog",
		}
		if err := validate.Struct(user); err != nil {
			fmt.Println("=== translate all ===")
			translateIndividual(trans, err)
		}
	}

	// ==============================
	fmt.Println("============================================================\n")
	// ==============================
	{
		user := &User{
			Name:    "gopher_forever",
			Level:   120,
			Email:   "dumpsterdiver.gmail.com",
			Tagline: "The quick brown fox jumps over the lazy dog",
		}
		if err := validate.Struct(user); err != nil {
			fmt.Println("=== translate all ===")
			translateOverride(trans, err)
		}
	}

}

// ================================================================================
func translateAll(trans ut.Translator, err error) {
	errs := err.(validator.ValidationErrors)
	terrs := errs.Translate(trans)
	for k := range terrs {
		fmt.Printf("%s: %s\n", k, terrs[k])
	}
	fmt.Println()
}

// ================================================================================
func translateIndividual(trans ut.Translator, err error) {
	errs := err.(validator.ValidationErrors)
	for _, e := range errs {
		fmt.Println(e.Translate(trans))
	}
	fmt.Println()
}

// ================================================================================
func translateOverride(trans ut.Translator, err error) {

	{
		registerFn := func(ut ut.Translator) error {
			return ut.Add("required", "{0} must have a value!", true)
		}

		translationFn := func(ut ut.Translator, fe validator.FieldError) string {
			t, _ := ut.T("required", fe.Field())
			return t
		}

		validate.RegisterTranslation("required", trans, registerFn, translationFn)
	}

	// ==============================
	errs := err.(validator.ValidationErrors)
	for _, e := range errs {
		fmt.Println(e.Translate(trans))
	}
	fmt.Println()
}

// ================================================================================
// noTranslator
func noTranslator(err error) {
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
