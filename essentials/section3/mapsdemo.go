package main

import (
	"fmt"
	"sort"
)

func main() {
	var capitals map[string]string = make(map[string]string)
	capitals["Afganistan"] = "Kabul"
	capitals["Canada"] = "Ottawa"
	capitals["Japan"] = "Tokyo"
	capitals["Kenya"] = "Nairobi"
	capitals["India"] = "New Delhi"
	capitals["Mexico"] = "Mexico City"
	capitals["South Korea"] = "Seoul"
	capitals["United Kingdom"] = "London"
	capitals["USA"] = "Washington D.C."
	capitals["Taiwan"] = "Taipei"

	// short hand notation to declare and initialize a map
	simpleCapitals := map[string]string{
		"Afganistan":     "Kabul",
		"Canada":         "Ottawa",
		"Japan":          "Tokyo",
		"Kenya":          "Nairobi",
		"India":          "New Delhi",
		"Mexico":         "Mexico City",
		"South Korea":    "Seoul",
		"United Kingdom": "London",
		"USA":            "Washington D.C.",
		"Taiwan":         "Taipei",
	}

	printMap(capitals)
	printSortedMap(simpleCapitals)
	crudOperations()

	// read https://blog.golang.org/go-maps-in-action to know about concurrency in using maps
}

func printMap(input map[string]string) {
	// Since map is an unordered data structure, the order of processing is unpredictable.
	fmt.Printf("%-25s\t%-25s\n", "Country", "Capital")
	for key, value := range input {
		fmt.Printf("%-25s\t%-25s\n", key, value)
	}
	fmt.Println()
}

func printSortedMap(input map[string]string) {
	// 1. there is no in-built way to get the keys of a map as a slice. so we manually create.
	// 2. sort the keys slice using sort package
	// 3. iterate keys slice
	// 3.1 print the value of keys
	keys := make([]string, len(input))

	i := 0
	for k := range input {
		keys[i] = k
		i++
	}
	// fmt.Println("unsorted keys:", keys)
	sort.Strings(keys)
	// fmt.Println("sorted keys:", keys)

	fmt.Printf("%-25s\t%-25s\n", "Country", "Capital")
	for _, k := range keys {
		fmt.Printf("%-25s\t%-25s\n", k, input[k])
	}
	fmt.Println()
}

func crudOperations() {
	fmt.Println("Creating an empty map")
	myMap := map[string]string{}
	fmt.Println("myMap", myMap, "\n")
	// add an element
	fmt.Println("Adding two items to map")
	myMap["fname"] = "Siva"
	myMap["lname"] = "Sankar"
	fmt.Println("myMap", myMap, "\n")

	fmt.Println("reading an item")
	lastName := myMap["lname"]
	fmt.Println("lname", lastName, "\n")

	fmt.Println("updating an item")
	myMap["lname"] = "Sankar R"
	fmt.Println("myMap", myMap, "\n")

	fmt.Println("deleting an item")
	delete(myMap, "lname")
	fmt.Println("myMap", myMap, "\n")

	fmt.Println("reading a non-existing item")
	fmt.Println("full name", myMap["fullname"], "\n")

	// testing the presence of a key in map
	// value can be replaced with _, to just check the presence. i.e., _, ok := myMap["key"]
	if value, ok := myMap["fname"]; ok {
		fmt.Println("fname exist. value:", value)
	} else {
		fmt.Println("fname doesn't exist. The default value is:", value)
	}

	if value, ok := myMap["lname"]; ok {
		fmt.Println("lname exist. value:", value)
	} else {
		fmt.Println("lname doesn't exist. The default value is:", value)
	}

}
