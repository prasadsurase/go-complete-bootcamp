package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// ---------------------------------------------------------
// EXERCISE: Unique Words
//
//  Create a program that prints the total and unique words
//  from an input stream.
//
//  1. Feed the shakespeare.txt to your program.
//
//  2. Scan the input using a new Scanner.
//
//  3. Configure the scanner to scan for the words.
//
//  4. Count the unique words using a map.
//
//  5. Print the total and unique words.
//
//
// EXPECTED OUTPUT
//
//  There are 99 words, 70 of them are unique.
//
// ---------------------------------------------------------

func main() {
	uniqWords := make(map[string]int)
	total := len(uniqWords)

	in := bufio.NewScanner(os.Stdin)

	for in.Scan() {
		words := strings.Fields(in.Text())
		for _, word := range words {
			total++
			uniqWords[word] = uniqWords[word] + 1
		}
	}
	fmt.Println("Total words:", total)
	fmt.Printf("\n %v", uniqWords)
}
