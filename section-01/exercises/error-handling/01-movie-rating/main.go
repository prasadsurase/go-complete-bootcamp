package main

import (
	"fmt"
	"os"
	"strconv"
)

//  1. Get the age from the command-line.
//
//  2. Return one of the following messages if the age is:
//     -> Above 17         : "R-Rated"
//     -> Between 13 and 17: "PG-13"
//     -> Below 13         : "PG-Rated"
//
// RESTRICTIONS:
//  1. If age data is wrong or absent let the user know.
//  2. Do not accept negative age.
//
// BONUS:
//  1. BONUS: Use if statements only twice throughout your program.
//  2. BONUS: Use an if statement only once.
//
// EXPECTED OUTPUT
//  go run main.go 18
//    R-Rated
//
//  go run main.go 17
//    PG-13
//
//  go run main.go 12
//    PG-Rated
//
//  go run main.go
//    Requires age
//
//  go run main.go -5
//    Wrong age: "-5"

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("Requires age\n")
		return
	}
	if age, err := strconv.Atoi(os.Args[1]); err != nil || age < 0 {
		fmt.Printf("Wrong age %q\n", os.Args[1])
	} else if age < 13 {
		fmt.Printf("PG-Rated\n")
	} else if age >= 13 && age <= 17 {
		fmt.Printf("PG-13\n")
	} else if age > 17 {
		fmt.Printf("R-Rated\n")
	}
}
