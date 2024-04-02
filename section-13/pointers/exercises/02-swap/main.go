// ---------------------------------------------------------
// EXERCISE: Swap
//
//  Using funcs, swap values through pointers, and swap
//  the addresses of the pointers.
//
//
//  1- Swap the values through a func
//
//     a- Declare two float variables
//
//     b- Declare a func that can swap the variables' values
//        through their memory addresses
//
//     c- Pass the variables' addresses to the func
//
//     d- Print the variables
//
//
//  2- Swap the addresses of the float pointers through a func
//
//     a- Declare two float pointer variables and,
//        assign them the addresses of float variables
//
//     b- Declare a func that can swap the addresses
//        of two float pointers
//
//     c- Pass the pointer variables to the func
//
//     d- Print the addresses, and values that are
//        pointed by the pointer variables
//
// ---------------------------------------------------------

package main

import "fmt"

func main() {
	var a, b float64

	a = 1.1
	b = 2.2
	fmt.Printf("Before swapping %f %f\n", a, b)
	swap(&a, &b)
	fmt.Printf("After swapping %f %f\n", a, b)

	var pA, pB *float64
	pA = &a
	pB = &b
	fmt.Printf("Before swapping %f %f\n", *pA, *pB)
	pA, pB = swapAddr(pA, pB)
	fmt.Printf("After swapping %f %f\n", *pA, *pB)
}

func swap(a, b *float64) {
	*a, *b = *b, *a
}

func swapAddr(a, b *float64) (*float64, *float64) {
	return b, a
}
