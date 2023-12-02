package main

import (
	"fmt"
	"os"
)

// ---------------------------------------------------------
// EXERCISE: Arg Count
//
//  1. Get arguments from command-line.
//  2. Print the expected outputs below depending on the number
//     of arguments.
//
// EXPECTED OUTPUT
//  go run main.go
//    Give me args
//
//  go run main.go hello
//    There is one: "hello"
//
//  go run main.go hi there
//    There are two: "hi there"
//
//  go run main.go I wanna be a gopher
//    There are 5 arguments
// ---------------------------------------------------------

func main() {
	if len(os.Args) == 1 {
		fmt.Println("Give me args\n")
	} else if len(os.Args) == 2 {
		fmt.Printf("There is one: %q\n", os.Args[1])
	} else if len(os.Args) == 3 {
		fmt.Printf("There are two: %s %s\n", os.Args[1], os.Args[2])
	} else {
		fmt.Printf("There are %d arguments\n", len(os.Args)-1)
	}
}
