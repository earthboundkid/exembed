package main

import (
	_ "embed"
	"fmt"
)

func main() {
	//go:embed file.txt
	var s string
	fmt.Print(s)
}
