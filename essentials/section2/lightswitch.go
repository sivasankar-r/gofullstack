package main

import "fmt"

var lightSwitchIsOn bool = false

// var lightSwitchIsOn = false //no error, you can do this
// lightSwitchIsOn := false // error, you cannot do this.

func main() {
	example1() // variable outside any function, var is mandatory
	// example2()
}

func example1() {
	printLightSwitchState()
	toggleLightSwitch()
	printLightSwitchState()
	toggleLightSwitch()
	printLightSwitchState()
}

func printLightSwitchState() {
	fmt.Println("The light switch is on: ", lightSwitchIsOn)
}

func toggleLightSwitch() {
	fmt.Println("toggling light switch...")
	lightSwitchIsOn = !lightSwitchIsOn
}
func example2() {
	isOn := false // inside a function, var is not required.
	fmt.Println("is on: ", isOn)
	fmt.Println("toggling switch...")
	isOn = !isOn
	fmt.Println("is on: ", isOn)
	fmt.Println("toggling switch...")
	isOn = !isOn
	fmt.Println("is on: ", isOn)
}
