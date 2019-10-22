package main

import "fmt"

func main() {
	i, j := 3, 55

	// The & operator generates a pointer to its operand
	iPtr := &i
	jPtr := &j

	fmt.Println("i:", i)       // print i
	fmt.Println("&i:", &i)     // print i's memory address
	fmt.Println("iPtr:", iPtr) // print iPtr which holds &i i.e., i's memory address

	// the * operator on a pointer variables dereferences and returns the underlying value of the pointer
	fmt.Println("Read i's value through its pointer using *iPtr, i:", *iPtr) // Read i's value through its pointer

	fmt.Println("Setting i's value through its pointer")
	*iPtr = 123                       // Set i's value through its pointer
	fmt.Println("new value of i:", i) // print i

	// fmt.Println("i:", i, "\tiPtr:", iPtr, "\t&i:", &i, "\t*iPtr:", *iPtr, "\t&iPtr:", &iPtr)

	*jPtr = *jPtr / 5 // divide j through the pointer
	fmt.Println(j)    // see the new value of j
}
