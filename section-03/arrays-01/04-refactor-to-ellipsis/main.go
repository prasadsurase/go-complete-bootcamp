package main

import (
	"fmt"
	"strings"
)

func main() {

	names := [...]string{"Sachin", "Nikhil", "Nilesh"}
	distances := [...]int{11, 22, 33, 44, 55}
	data := [...]uint8{'W', 'O', 'R', 'L', 'D'}
	ratios := [...]float64{3.14145}
	alives := [...]bool{true, false, true, false}
	zero := [...]uint8{}

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
