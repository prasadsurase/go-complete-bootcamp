package main

import (
	"fmt"
	"path"
)

func main() {
	dir, file := path.Split("go-complete-bootcamp/section-01/dirpathfinder.go")
	fmt.Println("directory: ", dir)
	fmt.Println("file: ", file)
}
