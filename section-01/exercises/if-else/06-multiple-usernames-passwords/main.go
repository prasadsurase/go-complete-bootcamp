package main

import (
	"fmt"
	"os"
)

// ---------------------------------------------------------
// CHALLENGE #2
//  Add one more user to the PassMe program below.
//
// EXAMPLE USERS
//  username: jack
//  password: 1888
//
//  username: inanc
//  password: 1879
//
// EXPECTED OUTPUT
//  go run main.go
//    Usage: [username] [password]
//
//  go run main.go hacker 42
//    Access denied for "hacker".
//
//  go run main.go jack 1888
//    Access granted to "jack".
//
//  go run main.go inanc 1879
//    Access granted to "inanc".
//
//  go run main.go jack 1879
//    Invalid password for "jack".
//
//  go run main.go inanc 1888
//    Invalid password for "inanc".
// ---------------------------------------------------------

const (
	usage        = "Usage: [username] [password]"
	errUser      = "Access denied for %q.\n"
	errPwd       = "Invalid password for %q.\n"
	accessOK     = "Access granted to %q.\n"
	user1, user2 = "jack", "inanc"
	pass1, pass2 = "1888", "1879"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println(usage)
		return
	}

	username, password := os.Args[1], os.Args[2]

	if username != user1 && username != user2 {
		fmt.Printf(errUser, username)
	} else if username == user1 && password == pass1 {
		fmt.Printf(accessOK, username)
	} else if username == user2 && password == pass2 {
		fmt.Printf(accessOK, username)
	} else {
		fmt.Printf(errPwd, username)
	}
}
