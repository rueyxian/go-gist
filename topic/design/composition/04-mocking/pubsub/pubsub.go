package pubsub

// Then writing an API, you shouldn't have to think about user testing and decoupling,
// the responsibility of mocking shouldn't be the person who write API, but the person who uses the API

type PubSub struct {
	host string
	// pretend there are more fields
}

func New(host string) *PubSub {
	ps := PubSub{
		host: host,
	}
	// pretend there is a specific implementation
	return &ps
}

func (ps *PubSub) Publish(key string, v interface{}) error {
	// pretend there is a specific implementation
	return nil
}

func (ps *PubSub) Subscribe(key string) error {
	// pretend there is a specific implementation
	return nil
}
