package main

import (
	"fmt"
	"net"
	"net/url"
)

func main() {
	s := "postgres://user:pass@host.com:5432/path?k=v#f"

	u, err := url.Parse(s)
	if err != nil {
		panic(err)
	}

	fmt.Println("scheme:", u.Scheme)

	fmt.Println("user:", u.User)

	p, _ := u.User.Password()
	fmt.Println("password:", p)

	fmt.Println("host:", u.Host)
	host, port, _ := net.SplitHostPort(u.Host)
	fmt.Println("host:", host)
	fmt.Println("port:", port)

	fmt.Println("path:", u.Path)
	fmt.Println("fragment:", u.Fragment)

	fmt.Println("raw_query:", u.RawQuery)
	m, _ := url.ParseQuery(u.RawQuery)
	fmt.Println("query:", m)
	fmt.Println("query [k][0]:", m["k"][0])
}
