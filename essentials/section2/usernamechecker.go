package main

import (
	"flag"
	"fmt"
	"log"
	"regexp"
)

// notice the `` in the regex
const usernameRegex string = `^@?(\w){1,15}$`

func main() {
	var input string
	flag.StringVar(&input, "username", "Gopher", "Username to check")
	flag.Parse()
	fmt.Println("Input username:", input, "\t Result: ", checkUsernameSyntax(input))
}

func checkUsernameSyntax(input string) bool {
	result := false

	r, err := regexp.Compile(usernameRegex)
	if err != nil {
		log.Fatal(err)
	}
	result = r.MatchString(input)
	return result
}
