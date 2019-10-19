package main

import "fmt"

func main() {

	// olympics LA onwards
	const (
		LosAngeles = 1984 + (iota * 4)
		Seoul
		Barcelona
		Atlanta
		Sydney
		Athens
		Beijing
		London
		Rio
		Tokyo
	)

	fmt.Println("These cities hosted or will host the summer olympics in the provided year...")
	fmt.Printf("%-18s %s \n", "City", "Year")
	fmt.Printf("%-18s %v \n", "Los Angeles", LosAngeles)
	fmt.Printf("%-18s %v \n", "Atlanta", Atlanta)
	fmt.Printf("%-18s %v \n", "Tokyo", Tokyo)
}
