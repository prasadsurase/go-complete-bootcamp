package main

import "fmt"

// ---------------------------------------------------------
// EXERCISE: Convert the strings
//
//   1. Loop over the words slice
//
//   2. In the loop:
//      1. Convert each string value to a byte slice
//      2. Print the byte slice
//      3. Append the byte slice to the `bwords`
//
//   3. Print the words using the `bwords`
//
// EXPECTED OUTPUT
//  [103 111 112 104 101 114]
//  [112 114 111 103 114 97 109 109 101 114]
//  [103 111 32 108 97 110 103 117 97 103 101]
//  [103 111 32 115 116 97 110 100 97 114 100 32 108 105 98 114 97 114 121]
//  gopher
//  programmer
//  go language
//  go standard library
// ---------------------------------------------------------

func main() {
	// Please uncomment the code below

	words := []string{
		"gopher",
		"programmer",
		"go language",
		"go standard library",
	}

	var bwords [][]byte

	for _, word := range words {
		wordBytes := []byte(word)
		fmt.Println(wordBytes)
		bwords = append(bwords, wordBytes)
	}

	for _, w := range bwords {
		fmt.Println(string(w))
	}
}
