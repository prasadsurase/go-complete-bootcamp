package main

import (
	"fmt"
	"strings"
)

func main() {
	msg := `
	
	The weather looks good.
I should go and play.



	`

	fmt.Println(msg)
	fmt.Println(strings.TrimSpace(msg))
}
