package main

import "fmt"

func main() {
	var s string = "Gopher"

	fmt.Println("First element:", s[0]) // this will print ascii value 71, instead of G
	fmt.Println("First element:", string(s[0]))
	fmt.Println("First element:", string("Gopher"[0])) // notice the use of actual literal
	fmt.Println("Last element:", string(s[len(s)-1]))  // notice the use of len function
	fmt.Println("1. Hello " + s + "!")
	fmt.Println("2. Hello", s, "!") // notice no space after hello
	fmt.Printf("3. Hello %v!\n", s)
}
