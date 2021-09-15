package main

import "sandbox/go-jottings/design-pattern/composition/04-mocking/pubsub"

// Go is a data oriented language, the best test should be also data oriented - concrete data in, concrete data out.
// Mocking should be our last resort, most of the time it can be solved by refactor our code to make it testable.

// What if mocking is really required?
// User of API can define their own set of interface that the API's statisfies
// Again, we discover interfaces, not design interfaces. Every developers can discover the interfaces as they need for themself

type publisher interface {
	Publish(key string, v interface{}) error
	Subscribe(key string) error
}

type mock struct{}

func (m *mock) Publish(key string, v interface{}) error {
	// mock for the publish
	return nil
}

func (m *mock) Subscribe(key string) error {
	// mock for the subscribe
	return nil
}

func main() {

	pubs := []publisher{
		pubsub.New("localhost"),
		&mock{},
	}

	for _, p := range pubs {
		p.Publish("key", "value")
		p.Subscribe("key")
	}
}
