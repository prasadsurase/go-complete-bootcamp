package main

import (
	"fmt"
	"strings"
)

func main() {
	books := [...]string{"Kafka's Revenge", "Stay Golden", "Everythingship"}

	upper := books
	lower := upper

	for i, name := range upper {
		upper[i] = strings.ToUpper(name)
	}

	for i, name := range lower {
		lower[i] = strings.ToLower(name)
	}

	fmt.Printf("books: %q\n", books)
	fmt.Printf("upper: %q\n", upper)
	fmt.Printf("lower: %q\n", lower)

}
