package main

import (
	"fmt"
	"unicode/utf8"
)

// ---------------------------------------------------------
// EXERCISE: Rune Manipulator
//
//  Please read the comments inside the following code.
//
// EXPECTED OUTPUT
//  Please run the solution.
// ---------------------------------------------------------

func main() {
	words := []string{
		"cool",
		"güzel",
		"jīntiān",
		"今天",
		"read 🤓",
	}

	_ = words

	for _, word := range words {
		fmt.Printf("\n--------------------------------------------------------------------")
		fmt.Printf("\n%q\n", word)
		// Print the byte and rune length of the strings
		// Hint: Use len and utf8.RuneCountInString
		fmt.Printf("Length: %d, Rune length: %d\n", len(word), utf8.RuneCountInString(word))

		// Print the bytes of the strings in hexadecimal
		// Hint: Use % x verb
		fmt.Printf("Bytes of strings in hex: % x\n", word)

		// Print the runes of the strings as rune literals
		// Hint: Use for range
		for _, c := range word {
			fmt.Printf("Bytes of strings in hex: % q\n", c)
		}

		// Print the first rune and its byte size of the strings
		// Hint: Use utf8.DecodeRuneInString
		r, size := utf8.DecodeRuneInString(word)
		fmt.Printf("First rune: %c, size: %d", r, size)

		// Print the last rune of the strings
		// Hint: Use utf8.DecodeLastRuneInString
		r, size = utf8.DecodeLastRuneInString(word)
		fmt.Printf("Last rune: %c, size: %d\n", r, size)

		// Slice and print the first two runes of the strings
		r1, s1 := utf8.DecodeRuneInString(word)
		r2, s2 := utf8.DecodeRuneInString(word[s1:])
		fmt.Printf("first 2 runes: %c %c, string: %s\n", r1, r2, word[:s1+s2])
		// Slice and print the last two runes of the strings

		// Convert the string to []rune
		// Print the first and last two runes
	}

}
