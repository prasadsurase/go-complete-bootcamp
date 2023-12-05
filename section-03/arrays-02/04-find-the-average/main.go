package main

import (
	"fmt"
	"os"
	"strconv"
)

//
//  Your goal is to fill an array with numbers and find the average.
//
//  1. Get the numbers from the command-line.
//
//  2. Create an array and assign the given numbers to that array.
//
//  3. Print the given numbers and their average.
//
// RESTRICTION
//   + Maximum 5 numbers can be provided
//   + If one of the arguments are not a valid number, skip it
//
// EXPECTED OUTPUT
//   go run main.go
//     Please tell me numbers (maximum 5 numbers).
//
//   go run main.go 1 2 3 4 5 6
//     Please tell me numbers (maximum 5 numbers).
//
//   go run main.go 1 2 3 4 5
//     Your numbers: [1 2 3 4 5]
//     Average: 3
//
//   go run main.go 1 a 2 b 3
//     Your numbers: [1 0 2 0 3]
//     Average: 2

func main() {
	if len(os.Args) > 6 {
		fmt.Println("Please tell me numbers (maximum 5 numbers).")
		return
	}
	numbers := [5]float64{}
	var sum float64
	for i, num := range os.Args[1:] {
		if val, err := strconv.ParseFloat(num, 64); err == nil {
			numbers[i] = val
			sum += float64(numbers[i])
		}
	}

	fmt.Printf("Numbers: %v\n", numbers)
	fmt.Printf("Sum: %.2f\n", sum)
	fmt.Printf("Average: %.2f\n", sum/float64(len(numbers)))
}
