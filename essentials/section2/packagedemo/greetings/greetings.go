package greetings

import f "fmt" // note f being used as alias for fmt pacakge

// PrintGreetings we indicate to Go that we want to export a funciton by upper casing the first letter in the function name
func PrintGreetings() {
	f.Println("in PrintGreetings()")
}

// This function is unexported, since it has a lowercase first letter in the function name
// Since this function is unexported, it will be visible only to the functions that are within the greetings package
func printGreetingsUnexported() {
	f.Println("in printGreetingsUnexported()")
}

func init() {
	MagicNumber = 1000
	f.Println("in greetings.go init(), MagicNumber:", MagicNumber)
}
