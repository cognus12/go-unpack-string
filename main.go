package main

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	str := "a4bc2d5e"
	fmt.Println("unpack string")
	res := unpack(str)
	fmt.Printf("Given: %v, output: %v", str, res)
}

func repeat(s string, n int) string {

	if len(s) > 1 {
		return ""
	}

	chars := make([]string, n)

	for i := 0; i < n-1; i++ {
		chars[i] = s
	}

	return strings.Join(chars[:], "")
}

func unpack(s string) string {

	result := ""

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
			nextIdx := i + 1

			if prevIdx < 0 || nextIdx > len(s)-1 {
				return ""
			}

			prev := string(s[prevIdx])
			next := rune(s[nextIdx])

			if unicode.IsDigit(next) {
				return ""
			}

			result += repeat(prev, n)
		} else {
			result += string(r)
		}
	}

	return result
}
