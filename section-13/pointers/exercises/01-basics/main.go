package main

import "fmt"

type computer struct {
	brand string
}

func main() {
	// create a nil pointer with the type of pointer to a computer

	var c *computer

	// compare the pointer variable to a nil value, and say it's nil
	if c == nil {
		fmt.Println("c is nil")
	} else {
		fmt.Println("c is not nil")
	}

	// create an apple computer by putting its address to a pointer variable
	a := &computer{brand: "Apple"}

	// put the apple into a new pointer variable
	b := a

	// compare the apples: if they are equal say so and print their addresses
	if a == b {
		fmt.Printf("Addresses are equal. %p %p\n", c, b)
	}

	// create a sony computer by putting its address to a new pointer variable
	s := &computer{brand: "Sony"}

	// compare apple to sony, if they are not equal say so and print their
	// addresses
	if c == s {
		fmt.Printf("Addresses are equal. %p %p\n", c, s)
	} else {
		fmt.Printf("Addresses are different. %p %p\n", c, s)
	}

	// put apple's value into a new ordinary variable
	apple := *c

	// print apple pointer variable's address, and the address it points to
	// and, print the new variable's addresses as well
	fmt.Printf("Addresses of pointer and its variable. %p %p\n", c, apple)

	// compare the value that is pointed by the apple and the new variable
	// if they are equal say so
	if apple == *c {
		fmt.Printf("Addresses are equal. %p %p\n", c, b)
	}

	// print the values:
	// the value that is pointed by the apple and the new variable
	fmt.Printf("Value: %v %v", apple, *c)

	// create a new function: change
	// the func can change the given computer's brand to another brand

	// change sony's brand to hp using the func — print sony's brand
	change(s, "HP")

	// write a func that returns the value that is pointed by the given *computer
	// print the returned value

	// write a new constructor func that returns a pointer to a computer
	// and call the func 3 times and print the returned values' addresses
}
func change(c *computer, brand string) {
	c.brand = brand
}
