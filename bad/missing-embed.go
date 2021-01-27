package main

import (
	"fmt"
)

//go:embed file.txt
var s string

func main() {
	fmt.Print(s)
}
