package main

import (
	"fmt"
	"reflect"

	"github.com/go-playground/validator/v10"
	"gopkg.in/guregu/null.v4"
)

type Test struct {
	ID        int64      `validate:"required,numeric"`
	Latitude  null.Float `validate:"required_with=Longitude,latitude"`
	Longitude null.Float `validate:"required_with=Latitude,longitude"`
}

func main() {
	validate := validator.New()

	validate.RegisterCustomTypeFunc(nullFloatValidator, null.Float{})

	b := &Test{
		ID: 1,
	}

	err := validate.Struct(b)
	fmt.Println(err)
}

func nullFloatValidator(field reflect.Value) interface{} {
	if valuer, ok := field.Interface().(null.Float); ok {
		if valuer.Valid {
			return valuer.Float64
		}
	}

	return nil
}
