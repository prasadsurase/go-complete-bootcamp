package main

import (
	"fmt"
	"os"
	"strconv"
)

//  go run main.go 3
//    X    0    1    2    3
//    0    0    0    0    0
//    1    0    1    2    3
//    2    0    2    4    6
//    3    0    3    6    9

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Give me the size of the table")
		return
	}

	max, err := strconv.Atoi(os.Args[1])
	if err != nil || max < 0 {
		fmt.Println("Wrong size")
		return
	}
	fmt.Printf("%5s", "X")

	for i := 0; i <= max; i++ {
		fmt.Printf("%5d", i)
	}
	fmt.Println("")
	for i := 0; i <= max; i++ {
		fmt.Printf("%5d", i)
		for j := 0; j <= max; j++ {
			fmt.Printf("%5d", i*j)
		}
		fmt.Println("")
	}
}
