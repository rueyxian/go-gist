package main

type Server struct {
	host string
	// pretend there are mor fields
}

func NewServer(host string) *Server {
	return &Server{host}
}

func (s *Server) Start() error {
	// pretend there is a specific implementation
	return nil
}

func (s *Server) Stop() error {
	// pretend there is a specific implementation
	return nil
}

func (s *Server) Wait() error {
	// pretend there is a specific implementation
	return nil
}

func main() {
	srv := NewServer("localhost")

	srv.Start()
	srv.Stop()
	srv.Wait()
}
