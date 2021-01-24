package main

type Server interface {
	Start() error
	Stop() error
	Wait() error
}

type server struct {
	host string
	// pretend there are more fields
}

func NewServer(host string) Server {
	return &server{host}
}

func (s *server) Start() error {
	// pretends there is a specific implementation
	return nil
}

func (s *server) Stop() error {
	// pretends there is a specific implementation
	return nil
}

func (s *server) Wait() error {
	// pretends there is a specific implementation
	return nil
}

func main() {

	srv := NewServer("localhost")

	srv.Start()
	srv.Stop()
	srv.Wait()
}
