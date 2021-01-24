package pubsub

type PubSub struct {
	host string
	//pretend there are more fields
}

func New(host string) *PubSub {
	// return &PubSub{
	//   host,
	// }

	ps := PubSub{
		host: host,
	}

	return &ps
}

func (ps *PubSub) Publish(k string, v interface{}) error {
	//pretend there are some implementations
	return nil
}

func (ps *PubSub) Subscribe(k string) error {
	//pretend there are some implementations
	return nil
}
