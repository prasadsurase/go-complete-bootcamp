package main

import (
	"fmt"
	"os"
)

func main() {
	// uncomment the code below
	// name := "and get the name from the command-line"

	// replace and concatenate the `name` variable
	// after `hi ` below

	msg := `hi ` + os.Args[1] + `!
how are you?`

	fmt.Println(msg)
}
