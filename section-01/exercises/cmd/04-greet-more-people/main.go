package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("There are", len(os.Args)-1, "people!")
	for i := 1; i < len(os.Args); i++ {
		fmt.Println("Hello great", os.Args[i], "!")
	}
	fmt.Println("Nice to meet you all.")
}
