package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
)

// ---------------------------------------------------------
// EXERCISE: Random Messages
//
//  Display a few different won and lost messages "randomly".
//
// HINTS
//  1. You can use a switch statement to do that.
//  2. Set its condition to the random number generator.
//  3. I would use a short switch.
//
// EXAMPLES
//  The Player wins: then randomly print one of these:
//
//  go run main.go 5
//    YOU WON
//  go run main.go 5
//    YOU'RE AWESOME
//
//  The Player loses: then randomly print one of these:
//
//  go run main.go 5
//    LOSER!
//  go run main.go 5
//    YOU LOST. TRY AGAIN?
// ---------------------------------------------------------

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Please provide a positive number between 1 and 10")
		return
	}

	guess, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println("%q is not a number", os.Args[1])
		return
	}
	if guess <= 0 {
		fmt.Println("%q is not a positive number", os.Args[1])
		return
	}

	// rand.Seed(time.Now().UnixNano())

	for itr := 1; itr <= 5; itr++ {
		if guess != rand.Intn(11) {
			continue
		}
		if itr == 1 {
			fmt.Println("Guessed right on first try!!!")
			return
		} else {
			switch rand.Intn(3) {
			case 0:
				fmt.Println("You won!")
			case 1:
				fmt.Println("You're awesome!")
			case 2:
				fmt.Println("perfect!")
			}
			return
		}
	}

	switch rand.Intn(2) {
	case 0:
		fmt.Println("Loser. Again?")
	case 1:
		fmt.Println("You lost. Try again?")
	}

}
