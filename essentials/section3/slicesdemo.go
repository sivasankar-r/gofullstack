package main

import "fmt"

func main() {

	// We use the make() built-in function to create a new slice of length 5.
	// Here we make a slice of integers of length 5.
	var mySlice []int = make([]int, 5)
	fmt.Printf("mySlice: %v\n\n", mySlice)

	updateSliceWithoutReturn(mySlice)
	fmt.Printf("After updateSliceWithoutReturn, mySlice: %v\n", mySlice)

	// We can use the len() built-in function to return the length of the slice
	fmt.Println("length of mySlice: ", len(mySlice))

	// We can use the cap() built-in function to return the capacity of the slice
	fmt.Println("capacity of mySlice: ", cap(mySlice), "\n")

	// Add a new element to mySlice, and notice the changes to the length and
	// capacity of the slice
	fmt.Println("After adding a new element to mySlice...\n")
	mySlice = append(mySlice, 1)
	fmt.Printf("Contents of mySlice: %v\n", mySlice)
	fmt.Println("length of mySlice: ", len(mySlice))
	fmt.Println("capacity of mySlice: ", cap(mySlice), "\n")

	// This short hand notation allows us to get the elements/to get the sub-slice from index 1 up to
	// (but not including) index 4.
	s := mySlice[1:4]
	fmt.Println("mySlice[1:4] ", s, "\n")

	// When you slice a slice, you get a reference back to that slice. Any changes you
	// make to the subslice will be reflected in the original slice.
	s[0] = 2 // this will cause mySlice[1] to be equal to 7
	fmt.Println("mySlice: ", mySlice)

	// All elements in myslice up to (not including) the element at index 4
	t := mySlice[:4]
	fmt.Println("mySlice[:4] ", t, "\n")

	// The elements from (and including) the element at index 1
	u := mySlice[1:]
	fmt.Println("mySlice[1:] ", u, "\n")

	slicesAppendDemo()
}

// Slices get passed by reference into functions, meaning, if make changes to
// a slice within a function, our changes will be reflected to the slice that
// was passed into the function.
func updateSliceWithoutReturn(input []int) {
	input[0] = 10
	input[1] = 9
	input[2] = 8
	input[3] = 7
	input[4] = 6
	fmt.Printf("in the updateSliceWithoutReturn(), slice: %v\n", input)
}

func slicesAppendDemo() {
	// Declaring and initializing a slice. Note it looks just like an array literal
	// except without the element count
	fmt.Println("rainbow:")
	rainbow := []string{"V", "I", "B", "G", "Y", "O"}
	fmt.Println(rainbow, "\n")

	// removing green color
	fmt.Println("removing green color...")
	rainbow = append(rainbow[:3], rainbow[4:]...)
	fmt.Println(rainbow, "\n")

	// JD joins the band in 2005
	fmt.Println("adding the missing red color at the end...")
	rainbow = append(rainbow, "R")
	fmt.Println(rainbow, "\n")

}
