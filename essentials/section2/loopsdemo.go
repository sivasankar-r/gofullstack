package main

import (
	"fmt"
	"time"
)

func main() {
	traditionalApproach()
	rangeApproach()
	singleCondition()
	infiniteLoop()
}

func traditionalApproach() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println("traditional approach, sum:", sum)
}

func rangeApproach() {
	/* create a slice */
	numbers := []int{1, 2, 3, 4, 5}
	sum := 0

	for _, value := range numbers {
		sum += value
	}
	fmt.Println("range approach, sum:", sum)
	fmt.Println()

	/* print the numbers */
	for i := range numbers {
		fmt.Println("Slice item", i, "is", numbers[i])
	}
	fmt.Println()

	/* create a map*/
	countryCapitalMap := map[string]string{"France": "Paris", "Italy": "Rome", "Japan": "Tokyo"}

	/* print map using keys*/
	for country := range countryCapitalMap {
		fmt.Println("Capital of", country, "is", countryCapitalMap[country])
	}

	/* print map using key-value*/
	for country, capital := range countryCapitalMap {
		fmt.Println("Capital of", country, "is", capital)
	}
	fmt.Println()
}

func singleCondition() {
	i := -4
	for i != 0 {
		fmt.Println("in single condition for loop, i:", i)
		i++
	}
	fmt.Println()
}

func infiniteLoop() {
	timer := time.NewTimer(time.Second * 9)
	for {
		fmt.Println("inside infinite loop")
		<-timer.C
		break
	}
	fmt.Println("completing now")
}
