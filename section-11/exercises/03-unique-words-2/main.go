package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

// ---------------------------------------------------------
// EXERCISE: Unique Words 2
//
//  Use your solution from the previous "Unique Words"
//  exercise.
//
//  Before adding the words to your map, remove the
//  punctuation characters and numbers from them.
//
//
// BE CAREFUL
//
//  Now the shakespeare.txt contains upper and lower
//  case letters too.
//
//
// EXPECTED OUTPUT
//
//  go run main.go < shakespeare.txt
//
//   There are 100 words, 69 of them are unique.
//
// ---------------------------------------------------------

func main() {
	// This is the regular expression pattern you need to use:
	// [^A-Za-z]+
	//
	// Matches to any character but upper case and lower case letters

	rgx := regexp.MustCompile(`[^A-Za-z]+`)

	uniqWords := make(map[string]int)
	total := len(uniqWords)

	in := bufio.NewScanner(os.Stdin)

	for in.Scan() {
		words := strings.Fields(in.Text())
		for _, word := range words {
			word = rgx.ReplaceAllString(word, "")
			total++

			uniqWords[strings.ToLower(word)] = uniqWords[strings.ToLower(word)] + 1
		}
	}
	fmt.Println("Total words:", total)
	fmt.Printf("\n %v", uniqWords)
}
