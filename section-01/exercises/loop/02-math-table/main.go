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
	if len(os.Args) < 3 {
		fmt.Println("Usage: [op=*/+-] [size]")
		return
	}

	optr := os.Args[1]
	max, err := strconv.Atoi(os.Args[2])
	if err != nil {
		fmt.Println("Wrong size. Error: ", err)
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
			var x int
			switch optr {
			case "*":
				x = i * j
			case "/":
				x = i / j
			case "+":
				x = i + j
			case "-":
				x = i - j
			}
			fmt.Printf("%5d", x)
		}
		fmt.Println("")
	}
}
