package main

type quacker interface {
	quack() string
}

// ================================================================================

type duck struct {
}

func (d duck) quack() string {
	return "quack quack!"
}

// ================================================================================

type chicken struct {
}

// ================================================================================

func main() {
	var _ quacker = duck{}
	// or ...
	// var _quacker = &duck{}     <-- same as var _quacker = new(duck)
	// or ...
	// var _quacker = (*duck)(nil)

	// but since chicken does not implements quacker, it will return error
	// var _ quacker = chicken{}
}
