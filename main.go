package main

import (
	"fmt"
	"strconv"
	"unicode"
)

func main() {
	str := "3abcd"
	fmt.Println("unpack string")
	res := unpack(str)
	fmt.Println(res)
}

func unpack(s string) string {

	if s == "" {
		return s
	}

	for i, r := range s {
		if unicode.IsDigit(r) {
			n, _ := strconv.Atoi(string(r))

			if n > 9 {
				return ""
			}

			prevIdx := i - 1

			if prevIdx < 0 {
				return ""
			}

			prev := string(s[i-1])
			fmt.Println(n, prev)
		}

	}

	return s
}
