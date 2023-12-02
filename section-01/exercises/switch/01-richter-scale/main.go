package main

import (
	"fmt"
	"os"
	"strconv"
)

// ---------------------------------------------------------
// MAGNITUDE                    DESCRIPTION
// ---------------------------------------------------------
// Less than 2.0                micro
// 2.0 and less than 3.0        very minor
// 3.0 and less than 4.0        minor
// 4.0 and less than 5.0        light
// 5.0 and less than 6.0        moderate
// 6.0 and less than 7.0        strong
// 7.0 and less than 8.0        major
// 8.0 and less than 10.0       great
// 10.0 or more                 massive

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Give me the magnitude of the earthquake")
		return
	}
	mag, err := strconv.ParseFloat(os.Args[1], 64)
	if err != nil {
		fmt.Printf("%q is not a number", os.Args[1])
		return
	}
	var desc string

	switch {
	case mag < 2:
		desc = "micro"
	case mag >= 2 && mag < 3:
		desc = "very minor"
	case mag >= 3 && mag < 4:
		desc = "minor"
	case mag >= 4 && mag < 5:
		desc = "light"
	case mag >= 5 && mag < 6:
		desc = "moderate"
	case mag >= 6 && mag < 7:
		desc = "strong"
	case mag >= 7 && mag < 8:
		desc = "major"
	case mag >= 8 && mag < 10:
		desc = "great"
	default:
		desc = "massive"
	}

	fmt.Printf("%.2f is %s\n", mag, desc)
}
