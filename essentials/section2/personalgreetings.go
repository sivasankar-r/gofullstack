package main

import (
	"flag"
	"fmt"
)

// steps to run:
// go run personalgreetings.go
// go run personalgreetings.go -help
// go run personalgreetings.go -name siva
// go run personalgreetings.go -name siva -msg "you are a rock star"
func main() {
	var name, msg string

	flag.StringVar(&name, "name", "Gopher", "Name of the Gopher")
	flag.StringVar(&msg, "msg", "Welcome!", "Welcome message")
	flag.Parse()
	fmt.Println("Hello", name, "!", msg)
}
