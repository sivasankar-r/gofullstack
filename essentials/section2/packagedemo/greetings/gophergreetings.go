package greetings

import f "fmt"

// MagicNumber will be initialized by the init() automatically
var MagicNumber int

// GopherGreetings func
func GopherGreetings() {
	f.Println("in GopherGreetings()")
	PrintGreetings()
	// we can call the below unexported function because this is within the same package
	printGreetingsUnexported()
}

// example of packages init() function
func init() {
	MagicNumber = 1729
	f.Println("in gophergreetings.go init(), MagicNumber:", MagicNumber)
}
