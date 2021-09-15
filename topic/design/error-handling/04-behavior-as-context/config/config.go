package config

import (
	"fmt"
	"reflect"
	"strings"
)

// ================================================================================

type Namer interface {
	Name() string
}

// ================================================================================

type InvalidParseError struct {
	Type reflect.Type
}

func (e *InvalidParseError) Error() string {
	if e.Type == nil {
		return "want: pointer struct | actual: nil"
	}
	if e.Type.Kind() != reflect.Ptr {
		return fmt.Sprintf("want: pointer struct | actual: value %s", e.Type.Kind())
	}
	if e.Type.Elem().Kind() != reflect.Struct {
		return fmt.Sprintf("want: pointer struct | actual: pointer %s", e.Type.Elem().Kind())
	}
	return fmt.Sprintf("want: pointer struct | actual: nil %s", e.Type.Elem().Kind())
}

func (e *InvalidParseError) Name() string { return "InvalidParseError" }

// ================================================================================
type InvalidArgumentError struct {
	Argument string
}

func (e *InvalidArgumentError) Error() string {
	return fmt.Sprintf("argument %s cannot be parsed", e.Argument)
}

func (e *InvalidArgumentError) Name() string { return "InvalidArgumentError" }

// ================================================================================

// just simulation, it didn't really do the job
func Parse(a interface{}, args []string) error {
	rv := reflect.ValueOf(a)
	rt := reflect.TypeOf(a)
	if rv.Kind() != reflect.Ptr || rv.Elem().Kind() != reflect.Struct {
		return &InvalidParseError{rt}
	}
	rv = rv.Elem()
	rt = rt.Elem()

	for i, arg := range args {
		if rt.Field(i).Name != extract(arg) {
			return &InvalidArgumentError{arg}
		}
	}
	return nil
}

func extract(s string) string {
	return strings.Split(s, "=")[0]
}

// ================================================================================
