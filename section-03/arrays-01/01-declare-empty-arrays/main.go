package main

import "fmt"

// names    : [3]string{"", "", ""}
//  distances: [5]int{0, 0, 0, 0, 0}
//  data     : [5]uint8{0x0, 0x0, 0x0, 0x0, 0x0}
//  ratios   : [1]float64{0}
//  alives   : [4]bool{false, false, false, false}
//  zero     : [0]uint8{}

func main() {

	var (
		names     [3]string
		distances [5]int
		data      [5]uint8
		ratios    [1]float64
		alives    [4]bool
		zero      [0]uint8
	)

	fmt.Printf("names     : %#v\n", names)
	fmt.Printf("distances : %#v\n", distances)
	fmt.Printf("data      : %#v\n", data)
	fmt.Printf("ratios    : %#v\n", ratios)
	fmt.Printf("alives    : %#v\n", alives)
	fmt.Printf("zero      : %#v\n", zero)

	fmt.Printf("names     : %T\n", names)
	fmt.Printf("distances : %T\n", distances)
	fmt.Printf("data      : %T\n", data)
	fmt.Printf("ratios    : %T\n", ratios)
	fmt.Printf("alives    : %T\n", alives)
	fmt.Printf("zero      : %T\n", zero)

	fmt.Printf("names     : %v\n", names)
	fmt.Printf("distances : %v\n", distances)
	fmt.Printf("data      : %v\n", data)
	fmt.Printf("ratios    : %v\n", ratios)
	fmt.Printf("alives    : %v\n", alives)
	fmt.Printf("zero      : %v\n", zero)

}
