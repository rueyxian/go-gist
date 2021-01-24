package barista

type brewer interface {
	brew() error
}

type Barista struct {
	name string
}

func New(name string) Barista {
	return Barista{
		name: name,
	}
}
