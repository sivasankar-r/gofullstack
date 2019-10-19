package main

import "fmt"

func main() {
	x := 2
	switch x {
	case 1:
		fmt.Println("x equals 1")
	case 2:
		fmt.Println("x equals 2")
	case 3:
		fmt.Println("x equals 3")
	default:
		fmt.Println("x is neither 1, 2 nor 3")
	}
}
