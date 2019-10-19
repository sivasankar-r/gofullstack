package main

import "fmt"

// iota is a predeclared identifier representing the untyped integer ordinal number of the current const specification in a (usually parenthesized) const declaration. It is zero-indexed.
const (
	_ = iota
	Red
	Green
	Yellow
)

func main() {
	fmt.Println("Red light code:", Red)
	fmt.Println("Green light code:", Green)
	fmt.Println("Yellow light code:", Yellow)
}
