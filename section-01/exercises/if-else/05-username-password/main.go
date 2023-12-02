package main

import (
	"fmt"
	"os"
)

// ---------------------------------------------------------
// CHALLENGE #1
//  Create a user/password protected program.
//
// EXAMPLE USER
//  username: jack
//  password: 1888
//
// EXPECTED OUTPUT
//  go run main.go
//    Usage: [username] [password]
//
//  go run main.go albert
//    Usage: [username] [password]
//
//  go run main.go hacker 42
//    Access denied for "hacker".
//
//  go run main.go jack 6475
//    Invalid password for "jack".
//
//  go run main.go jack 1888
//    Access granted to "jack".
// ---------------------------------------------------------

func main() {
	if len(os.Args) != 3 {
		fmt.Printf("Usage: [username] [password]")
		return
	}

	username, password := os.Args[1], os.Args[2]
	if username != "jack" {
		fmt.Printf("Access denied for %q.\n", username)
	} else if password != "1888" {
		fmt.Printf("Invalid password for %q.\n", username)
	} else {
		fmt.Printf("Access granted to %q.\n", username)
	}
}
