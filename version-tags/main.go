package main

import (
	"fmt"
	"strings"
)

var Version string = strings.TrimSpace(version)

func main() {
	fmt.Printf("Version %q\n", Version)
}
