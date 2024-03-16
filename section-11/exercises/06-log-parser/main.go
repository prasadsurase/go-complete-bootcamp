// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// ---------------------------------------------------------
// EXERCISE: Log Parser from Stratch
//
//  You've watched the lecture. Now, try to create the same
//  log parser program on your own. Do not look at the lecture,
//  and the existing source code.
//
//
// EXPECTED OUTPUT
//
//  go run main.go < log.txt
//
//   DOMAIN                             VISITS
//   ---------------------------------------------
//   blog.golang.org                        30
//   golang.org                             10
//   learngoprogramming.com                 20
//
//   TOTAL                                  60
//
//
//  go run main.go < log_err_missing.txt
//
//   wrong input: [golang.org] (line #3)
//
//
//  go run main.go < log_err_negative.txt
//
//   wrong input: "-100" (line #3)
//
//
//  go run main.go < log_err_str.txt
//
//   wrong input: "FOUR" (line #3)
//
// ---------------------------------------------------------

func main() {
	in := bufio.NewScanner(os.Stdin)

	domains := make(map[string]int)
	total := 0

	for in.Scan() {
		line := in.Text()
		fields := strings.Fields(line)
		if len(fields) < 2 {
			fmt.Printf("wrong input: %s\n", fields[0])
			return
		}
		val, err := strconv.Atoi(fields[1])
		if err != nil {
			fmt.Printf("wrong input: %s (line #%d)\n", fields[1], fields[0])
			return
		}
		if val < 0 {
			fmt.Printf("wrong input: %s (line #%d)\n", fields[1], fields[0])
			return
		}
		domains[fields[0]] = val
	}
	fmt.Printf("%25s %15s\n", "Domains", "Visits")
	fmt.Println(strings.Repeat("-", 50))
	for k, v := range domains {
		fmt.Printf("%25s %15d\n", k, v)
		total += v
	}
	fmt.Printf("%25s %15d\n", "Total", total)
}
