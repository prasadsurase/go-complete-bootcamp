package main

import (
	"fmt"
	"os"
	"strings"
)

// EXERCISE: Case Insensitive Search
//
//  Allow for case-insensitive searching
//
// EXAMPLE
//  Let's say that the user runs the program like this:
//    go run main.go LAZY
//
//  Or like this: go run main.go lAzY
//  Or like this: go run main.go lazy
//
//  For all cases above, the program should find
//  the "lazy" keyword.

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Please provide the search string.")
		return
	}
	sentence := "lazy cat jumps again and again and again"
	words := strings.Fields(sentence)

	// outer:
	for _, query_word := range os.Args[1:] {
		query_word = strings.ToLower(query_word)

	searcher:
		for j, word := range words {
			if query_word == word {
				fmt.Printf("#%2d: ", j+1)
				fmt.Println(word)
				continue searcher
			}
		}
	}
}
