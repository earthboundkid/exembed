package main

import (
	"bytes"
	_ "embed"
	"encoding/gob"
	"fmt"
)

var s = func() (s struct {
	Number   float64
	Weather  string
	Alphabet []string
}) {
	//go:embed value.gob
	var b []byte
	dec := gob.NewDecoder(bytes.NewReader(b))
	if err := dec.Decode(&s); err != nil {
		panic(err)
	}
	return
}()

func main() {
	fmt.Printf("s: %#v\n", s)
}
