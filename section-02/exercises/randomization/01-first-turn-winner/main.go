package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
)

// ---------------------------------------------------------
// EXERCISE: First Turn Winner
//
//  If the player wins on the first turn, then display
//  a special bonus message.
//
// RESTRICTION
//  The first picked random number by the computer should
//  match the player's guess.
//
// EXAMPLE SCENARIO
//  1. The player guesses 6
//  2. The computer picks a random number and it happens
//     to be 6
//  3. Terminate the game and print the bonus message
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
			fmt.Println("Guessed right!")
			return
		}
	}
	fmt.Println("You lost. Try again?")
}
