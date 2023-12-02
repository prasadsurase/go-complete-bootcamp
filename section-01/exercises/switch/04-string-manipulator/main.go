package main

import (
	"fmt"
	"os"
	"strings"
)

// EXERCISE: String Manipulator
//
//  1. Get the operation as the first argument.
//
//  2. Get the string to be manipulated as the 2nd argument.
//
// HINT
//
//	You can find the manipulation functions in the strings
//	package Go documentation (ToLower, ToUpper, Title).
//
// EXPECTED OUTPUT
//
//	go run main.go
//	  [command] [string]
//
//	  Available commands: lower, upper and title
//
//	go run main.go lower 'OMG!'
//	  omg!
//
//	go run main.go upper 'omg!'
//	  OMG!
//
//	go run main.go title "mr. charles darwin"
//	  Mr. Charles Darwin
//
//	go run main.go genius "mr. charles darwin"
//	  Unknown command: "genius"

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Usage: [command] [string]")
		fmt.Println("Available commands: lower, upper and title")
		return
	}

	command, msg := os.Args[1], os.Args[2]

	switch command {
	case "lower":
		msg = strings.ToLower(msg)
	case "upper":
		msg = strings.ToUpper(msg)
	case "title":
		msg = strings.Title(msg)
	default:
		fmt.Println("Unknown command %q", command)
		return
	}
	fmt.Println(msg)
}
