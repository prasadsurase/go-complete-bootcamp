package main

import (
	"fmt"
	"os"
	"strconv"
)

//  Your goal is to sort the given numbers from the command-line.
//
//  1. Get the numbers from the command-line.
//
//  2. Create an array and assign the given numbers to that array.
//
//  3. Sort the given numbers and print them.
//
// RESTRICTION
//   + Maximum 5 numbers can be provided
//   + If one of the arguments is not a valid number, skip it
//
// HINTS
//  + You can use the bubble-sort algorithm to sort the numbers.
//    Please watch this: https://youtu.be/nmhjrI-aW5o?t=7
//
//  + When swapping the elements, do not check for the last element.
//
//    Or, you will receive this error:
//    "panic: runtime error: index out of range"
//
// EXPECTED OUTPUT
//   go run main.go
//     Please give me up to 5 numbers.
//
//   go run main.go 6 5 4 3 2 1
//     Sorry. Go arrays are fixed. So, for now, I'm only supporting sorting 5 numbers...
//
//   go run main.go 5 4 3 2 1
//     [1 2 3 4 5]
//
//   go run main.go 5 4 a c 1
//     [0 0 1 4 5]

func main() {
	if len(os.Args) > 6 {
		fmt.Println("Please tell me numbers (maximum 5 numbers).")
		return
	}
	numbers := [5]float64{}

	for i, num := range os.Args[1:] {
		if val, err := strconv.ParseFloat(num, 64); err == nil {
			numbers[i] = val
		}
	}

	fmt.Printf("Numbers: %v\n", numbers)

	for i := 0; i < len(numbers)-1; i++ {
		for j := i + 1; j < len(numbers); j++ {
			if numbers[i] > numbers[j] {
				numbers[j], numbers[i] = numbers[i], numbers[j]
			}
		}
	}

	fmt.Printf("Sorted Numbers: %v\n", numbers)
}
