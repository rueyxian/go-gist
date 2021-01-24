package client

import "bufio"

type client struct {
	name   string
	reader *bufio.Reader
}

type temporary interface {
	Temporary() bool
}

func (c *client) TypeAsContext() {

}

func (c *client) BehaviorAsContext() {

}
