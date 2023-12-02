package main

import "fmt"

func main() {
	// EXERCISE: Convert and Fix #5
	//
	//  Fix the code.
	//
	// HINTS
	//   maximum of int8  can be 127
	//   maximum of int16 can be 32767
	//
	// EXPECTED OUTPUT
	//  1127

	// DO NOT TOUCH THESE VARIABLES
	min := int8(127)
	max := int16(1000)

	// FIX THE CODE HERE
	fmt.Println(max + int16(min))
}
