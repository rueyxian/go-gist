## func Indirect ¶

> func Indirect(v Value) Value

- Indirect returns the value that v points to. 
- If v is a nil pointer, Indirect returns a zero Value.
- If v is not a pointer, Indirect returns v.





## func (Value) Elem ¶

> func (v Value) Elem() Value

- Elem returns the value that the interface v contains or that the pointer v points to.
- It panics if v's Kind is not Interface or Ptr.
- It returns the zero Value if v is nil.


## differences
If a reflect.Value is a pointer, then v.Elem() is equivalent to reflect.Indirect(v). 
If it is not a pointer, then they are not equivalent:
- If the value is an interface then reflect.Indirect(v) will return the same value, while v.Elem() will return the contained dynamic value.
- If the value is something else, then v.Elem() will panic.

## source:
https://pkg.go.dev/reflect#Indirect
https://pkg.go.dev/reflect#Value.Elem
https://stackoverflow.com/questions/24318389/golang-elem-vs-indirect-in-the-reflect-package


