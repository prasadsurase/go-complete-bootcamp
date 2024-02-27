package main

import (
	"fmt"
	"os"
	"sort"
)

// ---------------------------------------------------------
// EXERCISE: Sort and write items to a file
//
//  1. Get arguments from command-line
//
//  2. Sort them
//
//  3. Write the sorted slice to a file
//
//
// EXPECTED OUTPUT
//
//   go run main.go
//     Send me some items and I will sort them
//
//   go run main.go orange banana apple
//
//   cat sorted.txt
//     apple
//     banana
//     orange
//
//
// HINTS
//
//   + REMEMBER: os.Args is a []string
//
//   + String slices are sortable using `sort.Strings`
//
//   + Use ioutil.WriteFile to write to a file.
//
//   + But you need to convert []string to []byte to be able to
//     write it to a file using the ioutil.WriteFile.
//
//   + To do that, create a new []byte and append the elements of your
//     []string.
//
// ---------------------------------------------------------

func main() {
	if len(os.Args) <= 1 {
		fmt.Println("Send me some items and I will sort them")
		return
	}

	items := os.Args[1:]
	fmt.Printf("Arguments: %v\n", items)
	sort.Strings(items)
	fmt.Printf("Sorted Arguments: %v\n", items)
	var data []byte
	for _, str := range items {
		data = append(data, str...)
		data = append(data, '\n')
	}

	err := os.WriteFile("sorted.txt", data, 0644)
	if err != nil {
		fmt.Printf("Error writing to file: %v", err)
		return
	}
}
