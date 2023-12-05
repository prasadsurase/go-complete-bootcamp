package main

import (
	"fmt"
	"strings"
)

func main() {
	var (
		names     [3]string
		distances [5]int
		data      [5]uint8
		ratios    [1]float64
		alives    [4]bool
		zero      [0]uint8
	)
	names[0] = "Sachin"
	names[1] = "Nikhil"
	names[2] = "Nilesh"

	distances[0] = 11
	distances[1] = 22
	distances[2] = 33
	distances[3] = 44
	distances[4] = 55

	data[0] = 'W'
	data[1] = 'O'
	data[2] = 'R'
	data[3] = 'L'
	data[4] = 'D'

	ratios[0] = 3.14145

	alives[0] = true
	alives[1] = false
	alives[2] = true
	alives[3] = false

	separator := "\n" + strings.Repeat("-", 50) + "\n"

	fmt.Println("\nnames", separator)
	for i := 0; i < len(names); i++ {
		fmt.Printf("names[%d]: %q\n", i, names[i])
	}

	fmt.Println("\ndistances", separator)
	for i := 0; i < len(distances); i++ {
		fmt.Printf("distances[%d]: %d\n", i, distances[i])
	}

	fmt.Println("\ndata", separator)
	for i := 0; i < len(data); i++ {
		fmt.Printf("data[%d]: %d\n", i, data[i])
	}

	fmt.Println("\nratios", separator)
	for i := 0; i < len(ratios); i++ {
		fmt.Printf("ratios[%d]: %f\n", i, ratios[i])
	}

	fmt.Println("\nalives", separator)
	for i := 0; i < len(alives); i++ {
		fmt.Printf("alives[%d]: %t\n", i, alives[i])
	}

	fmt.Println("\nzero", separator)
	for i := 0; i < len(zero); i++ {
		fmt.Printf("zero[%d]: %d\n", i, zero[i])
	}

	fmt.Println("\nnames", separator)
	for i, v := range names {
		fmt.Printf("names[%d]: %q\n", i, v)
	}

	fmt.Println("\ndistances", separator)
	for i, v := range distances {
		fmt.Printf("distances[%d]: %d\n", i, v)
	}

	fmt.Println("\ndata", separator)
	for i, v := range data {
		fmt.Printf("data[%d]: %d\n", i, v)
	}

	fmt.Println("\nratios", separator)
	for i, v := range ratios {
		fmt.Printf("ratios[%d]: %f\n", i, v)
	}

	fmt.Println("\nalives", separator)
	for i, v := range alives {
		fmt.Printf("alives[%d]: %t\n", i, v)
	}

	fmt.Println("\nzero", separator)
	for i, v := range zero {
		fmt.Printf("zero[%d]: %d\n", i, v)
	}
}
