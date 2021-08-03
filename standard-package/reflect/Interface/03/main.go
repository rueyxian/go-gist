package main

func main() {

}

// func main() {
//     var str []string
//     var v reflect.Value = reflect.ValueOf(&str)

//     v = v.Elem()

//     v = reflect.Append(v, reflect.ValueOf("a"))
//     v = reflect.Append(v, reflect.ValueOf("b"))
//     v = reflect.Append(v, reflect.ValueOf("c"), reflect.ValueOf("j, k, l"))

//     fmt.Println("Our value is a type of :", v.Kind())

//     vSlice := v.Slice(0, v.Len())
//     vSliceElems := vSlice.Interface()

//     fmt.Println("With the elements of : ", vSliceElems)
// }
