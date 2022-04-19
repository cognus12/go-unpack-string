package main

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	str := "aaa0bawd4a"
	res := unpack(str)
	fmt.Printf("Given: %v, output: %v", str, res)
}

func repeat(s string, n int) string {

	if n == 0 {
		return ""
	}

	if len(s) > 1 {
		return ""
	}

	return strings.Repeat(s, n-1)
}

func unpack(s string) string {

	if s == "" {
		return s
	}

	b := strings.Builder{}

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

			if n == 0 {

				current := b.String()
				trimmed := strings.TrimSuffix(current, prev)

				b.Reset()
				b.WriteString(trimmed)

				continue
			}

			b.WriteString(repeat(prev, n))
		} else {
			b.WriteString(string(r))
		}
	}

	return b.String()
}
