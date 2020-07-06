package main

import "fmt"

type Server struct {
	Host string
	Port int
}

type option func(*Server)

func Host(host string) option {
	return func(s *Server) {
		s.Host = host
	}
}

func Port(port int) option {
	return func(s *Server) {
		s.Port = port
	}
}

func NewServer(opts ...option) Server {
	s := Server{
		Host: "192.168.0.1",
		Port: 8080,
	}

	for _, opt := range opts {
		opt(&s)
	}

	return s
}

var host string
var port int

func main() {

	s1 := NewServer(Host("192.168.0.100"), Port(3000))
	s2 := NewServer(Port(3000))
	s3 := NewServer(Host("192.168.0.100"))

	fmt.Printf("%+v\n", s1)
	fmt.Printf("%+v\n", s2)
	fmt.Printf("%+v\n", s3)

}
