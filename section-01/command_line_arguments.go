package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Printf("%#v\n", os.Args)
	fmt.Println("Length: ", len(os.Args))
}
