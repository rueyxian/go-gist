

## Summary

| Method        | Code                                  | Argument      | CanSet() | Side Effect | Benchmarking<br />(composite type) |
| ------------- | ------------------------------------- | ------------- | -------- | ----------- | ---------------------------------- |
| ValueOf()     | ValueOf(T{})                          | value         | false    | -           | 3                                  |
|               | ValueOf(&T{})                         | pointer       | true     | yes         | 2                                  |
|               | ValueOf((\*T)(nil))                   | pointer (nil) | false    | -           | 1                                  |
| New(TypeOf()) | New(TypeOf(T{})).Elem()               | value         | true     | no          | 6                                  |
|               | New(TypeOf(&T{}).Elem()).Elem()       | pointer       | true     | no          | 5                                  |
|               | New(TypeOf((\*T)(nil)).Elem()).Elem() | pointer (nil) | true     | no          | 4                                  |

