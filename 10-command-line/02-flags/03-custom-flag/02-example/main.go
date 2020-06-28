package main

import (
	"flag"
	"fmt"
	"net/url"
)

type urlValue struct {
	url *url.URL
}

var u = &url.URL{}

func (v urlValue) String() string {
	if v.url != nil {
		return v.url.String()
	}
	return ""
}

func (v urlValue) Set(s string) error {

	if u, err := url.Parse(s); err == nil {
		*v.url = *u
	} else {
		return err
	}
	return nil
}

func init() {

	u, _ = url.Parse("https://golang.org/pkg/net/url/")
	flag.Var(&urlValue{u}, "url", "nothing special")
	flag.Parse()
}

func main() {

	fmt.Printf("scheme: %v \n", u.Scheme)
	fmt.Printf("  host: %v \n", u.Host)
	fmt.Printf("  path: %v \n", u.Path)

}
