package main

import (
	"bytes"
	"encoding/gob"
	"io"
	"math"
	"net/http"
	"os"
)

func main() {
	s := struct {
		Number   float64
		Weather  string
		Alphabet []string
	}{
		Number:   math.Pi,
		Weather:  getString(),
		Alphabet: getSlice(),
	}

	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	must(enc.Encode(&s))
	must(os.WriteFile("value.gob", buf.Bytes(), os.ModePerm))
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}

func getString() string {
	rsp, err := http.Get("http://wttr.in?format=1")
	must(err)
	defer rsp.Body.Close()
	b, err := io.ReadAll(rsp.Body)
	must(err)
	return string(b)
}

func getSlice() []string {
	var s []string
	for c := 'A'; c <= 'Z'; c++ {
		s = append(s, string(c))
	}
	return s
}
