//go:build prod
// +build prod

package main

import (
	_ "embed"
)

//go:embed version.txt
var version string
