package main

import (
	// _ "embed"
	"fmt"
)

func main() {
	//go:embed quine.go
	var src string
	fmt.Print(src)
}
