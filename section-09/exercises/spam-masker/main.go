package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]
	if len(args) != 1 {
		fmt.Println("Give something to mask")
		return
	}

	const (
		mask = '*'
	)

	var (
		text      = args[0]
		size      = len(text)
		buff      = make([]byte, 0, size) //initial length should be 0
		mapper    = "http://"
		mapperLen = len(mapper)
		in        = false
	)

	// buff = append(buff, text...)

	for indx := 0; indx < size; indx++ {
		if len(text[indx:]) >= mapperLen && text[indx:indx+mapperLen] == mapper {
			in = true
			buff = append(buff, mapper...)
			indx += mapperLen
		}
		c := text[indx]

		switch c {
		case ' ', '\t', '\n':
			in = false
		}

		if in {
			c = mask
		}
		buff = append(buff, c)
	}
	fmt.Println(string(buff))
}
