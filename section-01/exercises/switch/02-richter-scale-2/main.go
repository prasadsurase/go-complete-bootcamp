package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Tell me the magnitude of the earthquake in human terms.")
		return
	}

	mag := os.Args[1]
	var desc string

	switch mag {
	case "micro":
		desc = "less than 2.0"
	case "very minor":
		desc = "2 - 2.9"
	case "minor":
		desc = "3 - 3.9"
	case "light":
		desc = "4 - 4.9"
	case "moderate":
		desc = "5 - 5.9"
	case "strong":
		desc = "6 - 6.9"
	case "major":
		desc = "7 - 7.9"
	case "great":
		desc = "8 - 9.9"
	case "massive":
		desc = "10+"
	default:
		fmt.Printf("%q's richter scale is unknown\n", mag)
		return
	}

	fmt.Printf("%s's richter scale is %s\n", mag, desc)
}
