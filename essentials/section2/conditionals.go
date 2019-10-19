package main

import "fmt"

func main() {
	x := 9

	// if
	if x == 9 {
		fmt.Println("x == 9")
	}

	// if else
	if y := 4; y == 5 {
		fmt.Println("y == 5")
	} else {
		fmt.Println("y != 5")
	}

	// if else-if else
	z := 5
	if z == 1 {
		fmt.Println("z == 1")
	} else if z == 2 {
		fmt.Println("z == 2")
	} else if z == 3 {
		fmt.Println("z == 3")
	} else {
		fmt.Println("z does not equal to 1, 2, 3")
	}
}
