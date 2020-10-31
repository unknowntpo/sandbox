package main

import (
	"fmt"
	"net"
	"net/url"
)

func showResult(s string) {
	u, err := url.Parse(s)
	if err != nil {
		panic(err)
	}

	fmt.Println("Input URL: ", s)

	fmt.Println("Scheme: ", u.Scheme)

	fmt.Println("Username: ", u.User.Username())

	p, _ := u.User.Password()
	fmt.Println("Password: ", p)

	fmt.Println("Host: ", u.Host)
	host, port, _ := net.SplitHostPort(u.Host)
	fmt.Println("\tHost: ", host)
	fmt.Println("\tPort: ", port)

	fmt.Println("Path:", u.Path)
	fmt.Println("Fragment:", u.Fragment)

	fmt.Println("RawQuery: ", u.RawQuery)
	m, _ := url.ParseQuery(u.RawQuery)
	fmt.Println(m)
	fmt.Printf("Data type of RawQuery: %T\n", m)
}
func main() {
	urls := []string{
		"postgres://user:pass@host.com:5432/path?k=v#f",
		"http://www.example.com/path?key=value#f",
	}
	for _, url := range urls {
		showResult(url)
		fmt.Println("-----------------------------------------")
	}
}
