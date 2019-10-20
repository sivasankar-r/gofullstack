package main

import "fmt"

func main() {
	// defining the size is mandatory in case of array declaration
	var myArray [3]int

	// when not initialized with values, it will default to their zero values.
	// In this case the default value is 0
	fmt.Printf("myArray: %v\n\n", myArray)

	// Arrays are passed by value to functions, meaning a copy of the array is passed
	// and not a reference (pointer) to the array.
	updateArrayWithoutReturn(myArray)
	fmt.Printf("After updateArrayWithoutReturn(), myArray: %v\n\n", myArray)

	myArray = updateArrayWithReturn(myArray)
	fmt.Printf("After updateArrayWithReturn, myArray: %v\n\n", myArray)

	// len(inputArray) returns the length of the array
	fmt.Println("Length of myArray: ", len(myArray), "\n")

	// arrays use zero-based index
	fmt.Println("Second element in the array", myArray[1], "\n")

	// arrays of different sizes are considered as different size, though they use the same data type.

	var aArray [5]int
	// the below line will not compile, since myArray is [3]int and aArray is [5]int. they are incompatible
	// aArray = myArray

	// One-liner to declare and initiallize an array using the := operator. do not use var when using :=
	bArray := [5]int{1, 2, 3, 4, 5}
	aArray = bArray
	fmt.Println("aArray:", aArray)
	fmt.Println("bArray:", bArray, "\n")

	// when an array is assigned to another, then all the values are copied. so both of them are different arrays in memory
	fmt.Printf("aArray memory address: %p\n", &aArray)
	fmt.Printf("bArray memory address: %p\n\n", &bArray)

	// array equality comparing is done by verifying the equality of values in each index but not the memory address of the arrays
	if aArray == bArray {
		fmt.Println("values in aArray and bArray are equal\n")
	} else {
		fmt.Println("values in aArray and bArray are not equal\n")
	}

	twoDimensional()
}

// since arrays are passed by value, the updates on the input array will not affect the original array
func updateArrayWithoutReturn(input [3]int) {
	input[0] = 3
	input[1] = 6
	input[2] = 9
	fmt.Printf("in the updateArrayWithoutReturn(), array: %v\n", input)
}

// 2nd attempt at populating the integer array
func updateArrayWithReturn(input [3]int) [3]int {
	input[0] = 3
	input[1] = 6
	input[2] = 9
	fmt.Printf("in the updateArrayWithReturn(), array: %v\n", input)
	return input
}

func twoDimensional() {
	var twoD [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("two dimensional array:", twoD, "\n")
}
