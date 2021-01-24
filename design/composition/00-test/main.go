package main

import "sandbox/go-jottings/design-pattern/composition/00-test/pubsub"

type publisher interface {
	Publish(v string, k interface{}) error
	Subscribe(v string) error
}

type mock struct{}

func (m *mock) Publish(k string, v interface{}) error {
	//some specific implementation
	return nil
}

func (m *mock) Subscribe(k string) error {
	//some specific implementation
	return nil
}

func main() {

	pubs := []publisher{
		pubsub.New("1234"),
		&mock{},
	}

	for _, p := range pubs {
		p.Publish("Key", "Value")
		p.Subscribe("Key")
	}

}
