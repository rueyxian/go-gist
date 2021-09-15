package main

// ================================================================================

type lamer interface {
	lame() string
}

// ================================================================================

type dog struct{}

func (_ dog) lame() string {
	return "I'm so lame"
}

// ================================================================================

type cat string

func (_ cat) lame() string {
	return "I'm freaking awesome"
}

// ================================================================================

type gopher int

func (_ gopher) awesome() string {
	return "I'm awesome too"
}

// ================================================================================

func main() {

	// struct value
	var _ lamer = dog{}

	// pointer to struct
	var _ lamer = (*dog)(nil)

	// simple type value
	var _ lamer = cat("")

	// pointer to simple type
	var _ lamer = (*cat)(nil)

	// compile error
	var _ lamer = gopher(0)

}
