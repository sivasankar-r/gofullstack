package main

import (
	"fmt"

	"github.com/sivasankar-r/gofullstack/essentials/section2/packagedemo/greetings"
)

func main() {
	greetings.PrintGreetings()  // this is in ../greetings/greeting.go file
	greetings.GopherGreetings() // this is in ../greetings/gophergreetings.go file

	fmt.Println("Magic Number", greetings.MagicNumber)
}
