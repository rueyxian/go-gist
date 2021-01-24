package dsv

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"reflect"
)

// ==============================

type NotPointerError struct {
	Type reflect.Type
}

func (e *NotPointerError) Error() string {
	return fmt.Sprintf("dsv: %v has to be pointer", e.Type.String())
}

// ==============================

type NotSliceError struct {
	Type reflect.Type
}

func (e *NotSliceError) Error() string {
	return fmt.Sprintf("dsv: %v has to be slice", e.Type.String())
}

// ==============================

func Unmarshal(data []byte, v interface{}) error {
	rv := reflect.ValueOf(v)
	// if rv.Kind() != reflect.Ptr {
	//   return &NotPointerError{reflect.TypeOf(v)}
	// }

	if rv.Kind() != reflect.Slice {
		return &NotSliceError{reflect.TypeOf(v)}
	}

	r := bufio.NewReader(bytes.NewReader(data))

	for {
		line, err := r.ReadString('\n')

		if err != nil {

			switch e := err.(type) {

			default:
				if err == io.EOF {
					return
				}

			}

			fmt.Println(string(line))

		}

	}

	return nil

}
