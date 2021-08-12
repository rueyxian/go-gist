package main

import (
	"fmt"
	"math/rand"
	"reflect"
	"runtime"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {

	wrapTimed := func(a interface{}) interface{} {
		rv := reflect.ValueOf(a)
		if rv.Kind() != reflect.Func {
			panic("expect: kind is func")
		}

		fn := func(in []reflect.Value) []reflect.Value {
			start := time.Now()
			out := rv.Call(in)
			end := time.Now()
			fmt.Printf("%s: %v\n", runtime.FuncForPC(rv.Pointer()).Name(), end.Sub(start))
			return out
		}

		wrapper := reflect.MakeFunc(rv.Type(), fn)
		return wrapper.Interface()
	}

	a := wrapTimed(taskA).(func())
	b := wrapTimed(taskB).(func(string) error)

	a()
	b("troll")

}

func taskA() {
	time.Sleep(time.Duration(rand.Intn(2000)) * time.Millisecond)
}

func taskB(itdoesnotmatter string) error {
	time.Sleep(time.Duration(rand.Intn(4000)) * time.Millisecond)
	_ = itdoesnotmatter
	return nil
}
