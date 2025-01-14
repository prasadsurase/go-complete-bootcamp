package main

import "fmt"

// ---------------------------------------------------------
// EXERCISE: Print the runes
//
//  1. Loop over the "console" word and print its runes one by one,
//     in decimals, hexadecimals and binary.
//
//  2. Manually put the runes of the "console" word to a byte slice, one by one.
//
//     As the elements of the byte slice use only the rune literals.
//
//     Print the byte slice.
//
//  3. Repeat the step 2 but this time, as the elements of the byte slice,
//     use only decimal numbers.
//
//  4. Repeat the step 2 but this time, as the elements of the byte slice,
//     use only hexadecimal numbers.
//
//
// EXPECTED OUTPUT
//   Run the solution to see the expected output.
// ---------------------------------------------------------

func main() {
	const word = "console"
	bw := []byte(word)
	fmt.Printf("%15s %15s %15s %15s\n", "Rune", "Decimal", "Hexadecimal", "Binary")
	for _, b := range bw {
		fmt.Printf("%15c %15d %15x %15b\n", b, b, b, b)
	}

	// print the word manually using runes
	fmt.Printf("with runes: %s\n", string([]byte{'c', 'o', 'n', 's', 'o', 'l', 'e'}))

	// print the word manually using decimals
	fmt.Printf("with decimals: %s\n", string([]byte{99, 111, 110, 115, 111, 108, 101}))

	// print the word manually using hexadecimals
	fmt.Printf("with hexadecimals: %s\n", string([]byte{0x63, 0x6f, 0x6e, 0x73, 0x6f, 0x6c, 0x65}))
}
