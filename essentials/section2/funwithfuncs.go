package main

import "fmt"

var y = 4

// a function that takes one input and returns nothing
func oddOrEven(x int) {
	// the scope of variables in a function is local i.e., the variable x is accessible only within the function
	if x%2 == 0 {
		fmt.Println("x:", x, "is even")
	} else {
		fmt.Println("x:", x, "is odd")
	}

	// the variable y is visible inside the function, because it exists in the global scope
	if y%2 == 0 {
		fmt.Println("y:", y, "is even")
	} else {
		fmt.Println("y:", y, "is odd")
	}
}

// a function with two input params and return an output of int type
func sumOfTwoInts(x int, y int) int {
	return x + y
}

// a function that returns multiple named output parameters. Notice for same input type parameters you can club together
// same thing can be done for output also like (sum, diff int)
func sumAndDiff(x, y int) (sum int, diff int) {
	return x + y, x - y
}

// function with unnamed output parameters
func sumAndDiffUnnamed(x, y int) (int, int) {
	return x + y, x - y
}

func multiSum(args ...int) int {
	sum := 0
	// using range approach
	for _, value := range args {
		sum += value
	}

	// traditional approach
	// for i := 0; i < len(args); i++ {
	// 	sum += args[i]
	// }
	return sum
}

func main() {
	oddOrEven(23)
	fmt.Println("sum of 12 and 3:", sumOfTwoInts(12, 3))
	s, d := sumAndDiff(12, 3)
	fmt.Println("Input: 12 and 3, Sum:", s, " Diff:", d)
	s1, d1 := sumAndDiffUnnamed(2, 3)
	fmt.Println("Input: 2 and 3, Sum:", s1, " Diff:", d1)
	fmt.Println("Multisum result: ", multiSum(1, 2, 3, 4, 5))

	// example of anonymous function
	func() {
		fmt.Println("anonymous function")
	}()

	assigned := func() {
		fmt.Println("assigned anonymous function")
	}
	assigned()
}
