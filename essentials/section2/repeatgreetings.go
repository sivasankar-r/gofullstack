package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) > 1 {
		repeatCount, err := strconv.Atoi(os.Args[1])
		if err != nil {
			log.Fatal(err)
		} else {
			for i := 0; i < repeatCount; i++ {
				fmt.Println("Hello Gopher!")
			}
		}
	} else {
		log.Fatal("Pass an integer argument to the program.")
	}
}
