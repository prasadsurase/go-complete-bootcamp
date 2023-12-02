package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Please provide min and max numbers")
		return
	}

	min, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Printf("%q is not a number", min)
		return
	}
	max, err := strconv.Atoi(os.Args[2])
	if err != nil {
		fmt.Printf("%q is not a number", max)
		return
	}

	if min > max {
		min, max = max, min
	}

	sum := 0
	for i := min; i <= max; i++ {
		if i%2 == 0 {
			sum += i
			fmt.Printf("%d", i)
			if i == max || i+1 == max {
				fmt.Print(" = ")
			} else {
				fmt.Print(" + ")
			}
		}
	}
	fmt.Printf("%d\n", sum)
}
