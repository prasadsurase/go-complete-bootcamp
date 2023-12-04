package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
)

//  Let the player guess two numbers instead of one.
//
// HINT:
//  Generate random numbers using the greatest number
//  among the guessed numbers.
//
// EXAMPLES
//  go run main.go 5 6
//  Player wins if the random number is either 5 or 6.

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Please provide a positive number between 1 and 10")
		return
	}

	guess1, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Printf("%q is not a number\n", os.Args[1])
		return
	}
	if guess1 <= 0 {
		fmt.Printf("%q is not a positive number\n", os.Args[1])
		return
	}

	guess2, err := strconv.Atoi(os.Args[2])
	if err != nil {
		fmt.Printf("%q is not a number\n", os.Args[2])
		return
	}
	if guess2 <= 0 {
		fmt.Printf("%q is not a positive number\n", os.Args[2])
		return
	}

	for itr := 1; itr <= 5; itr++ {
		num := rand.Intn(11)
		fmt.Println("Random number:", num)
		fmt.Println("guess 1:", guess1)
		fmt.Println("guess 2:", guess2)
		if guess1 != num && guess2 != num {
			continue
		}
		if itr == 1 {
			fmt.Println("Guessed right on first try!!!")
			return
		} else {
			fmt.Println("Guessed right!")
			return
		}
	}
	fmt.Println("You lost. Try again?")
}
