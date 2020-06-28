package main

import (
	"flag"
	"fmt"
	"net/url"
)

type URLValue struct {
	URL *url.URL
}

func (v URLValue) String() string {
	if v.URL != nil {
		return v.URL.String()
	}
	return ""
}

func (v URLValue) Set(s string) error {
	if u, err := url.Parse(s); err != nil {
		return err
	} else {
		*v.URL = *u
	}
	return nil
}

var u = &url.URL{}

func init() {
	// u, _ = url.Parse("https://golang.org/pkg/flag/")
	// flag.Var(&URLValue{u}, "url", "URL to parse")

	// flag.Parse()
}

func main() {
	fs := flag.NewFlagSet("ExampleValue", flag.ExitOnError)
	fs.Var(&URLValue{u}, "url", "URL to parse")
	// fs.Parse([]string{"-url", "https://golang.org/pkg/flag/"})

	flag.Parse()

	fmt.Printf("scheme: %v \n", u.Scheme)
	fmt.Printf("  host: %v \n", u.Host)
	fmt.Printf("  path: %v \n", u.Path)

}
