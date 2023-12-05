package main

import (
	"fmt"
	"strings"
)

//
//   In this exercise, your goal is to display a few famous scientists
//   in a pretty table.
//
//   1. Create a multi-dimensional array
//   2. In each inner array, store the scientist's name, lastname and his/her
//      nickname
//   3. Print their information in a pretty table using a loop.
//
// EXPECTED OUTPUT
//   First Name      Last Name       Nickname
//   ==================================================
//   Albert          Einstein        time
//   Isaac           Newton          apple
//   Stephen         Hawking         blackhole
//   Marie           Curie           radium
//   Charles         Darwin          fittest

func main() {
	persons := [...][3]string{
		{"Albert", "Einstein", "time"},
		{"Isaac", "Newton", "apple"},
		{"Stephan", "Hawking", "blackhole"},
		{"Marie", "Curie", "radium"},
		{"Charles", "Darwin", "fittest"},
	}

	fmt.Printf("%-15v %-15v %-15v\n", "Firstname", "Lastname", "Nickname")
	fmt.Println("\n" + strings.Repeat("=", 45))
	for _, person := range persons {
		fmt.Printf("%-15v %-15v %-15v\n", person[0], person[1], person[2])
	}
}
