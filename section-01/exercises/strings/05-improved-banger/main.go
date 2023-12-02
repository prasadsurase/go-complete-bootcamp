package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	msg := os.Args[1]
	fmt.Println(msg + strings.Repeat("!", len(msg)))
}
