package main

import (
	"fmt"
	"go-unpack-string/pkg/unpackstring"
)

func main() {
	str := `qwe\45`
	res := unpackstring.Unpack(str)
	fmt.Printf("Given: %v, output: %v", str, res)
}
