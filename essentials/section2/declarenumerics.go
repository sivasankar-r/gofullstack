package main

import (
	"fmt"
	"reflect"
)

func main() {
	uninitializedInt32()
	sumOfTwoFloats()
	autoDetect()
	groupingConstants()
	groupingVariables()
}

func uninitializedInt32() {
	// Declare a 32 bit integer without initializing. variables declared without an explicit initial value are given their zero value
	// int is not alias for int32. both are different types but 32 bits
	var myInteger int32
	fmt.Println("value of myInteger: ", myInteger)
	fmt.Println()
}

func sumOfTwoFloats() {
	// declare and initialize two float64 and then show their sum
	var myFloat float64 = 36.333
	var anotherFloat float64 = 306.96969696
	fmt.Println("sum of floats: ", myFloat+anotherFloat)
	fmt.Println()
}

func autoDetect() {
	// go will figure out what the type is
	x, y, z := 0, 1, 2
	// notice the printf instead of println
	fmt.Printf("x: %d\ty: %d\tz: %d\n", x, y, z)
	fmt.Println("Type of x: ", reflect.TypeOf(x), " y: ", reflect.TypeOf(y), " z: ", reflect.TypeOf(z))
	fmt.Println()

	// note the use of fmt.Sprintf and %T to get the type. However reflect package is more efficient
	// read more at https://stackoverflow.com/a/27160765/1034051
	myFloat := 306.96969696
	fmt.Println("myFloat: ", myFloat)
	fmt.Println("type of myFloat: ", fmt.Sprintf("%T", myFloat))
	fmt.Println()

	myComplex := 5 + 24i
	fmt.Println("myComplex: ", myComplex)
	fmt.Println("type of myComplex: ", reflect.TypeOf(myComplex))
	fmt.Println()
}

func groupingConstants() {
	const (
		speedOfLight = 299792458 // units m/sec
		pi           = 3.14
		favNumber    = 1729 //ramanujanNumber
	)
	// %v prints the variable in the default format
	fmt.Printf("speed of light: %v\tpi: %v\tfavNumber: %v\n", speedOfLight, pi, favNumber)
	fmt.Println()
	/*
	   The default format for %v is:

	   bool:                    %t
	   int, int8 etc.:          %d
	   uint, uint8 etc.:        %d, %#x if printed with %#v
	   float32, complex64, etc: %g
	   string:                  %s
	   chan:                    %p
	   pointer:                 %p

	   Read more formatting details at https://golang.org/pkg/fmt
	*/
}

func groupingVariables() {
	// example of grouping variable declarations/initializations
	var (
		a int = 0        //an integer
		b     = 1.8 + 3i //complex number
		c     = 2.7      // floating point number
	)
	fmt.Printf("a: %v\tb: %v\tc: %v\n\n", a, b, c)
}
