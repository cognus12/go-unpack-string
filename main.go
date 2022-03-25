package main

import (
	"fmt"
	"strconv"
	"unicode"
)

func main() {
	str := "a4bc2d5e"
	fmt.Println("unpack string")
	res := unpack(str)
	fmt.Println(res)
}

func unpack(s string) string {

	for i, r := range s {
		if unicode.IsDigit(r) {
			n, _ := strconv.Atoi(string(r))
			prev := string(s[i-1])
			fmt.Println(n, prev)
		}

	}

	return s
}
