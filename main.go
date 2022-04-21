package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	str := `qwe\4\5`
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

func checkSymbol(r, prev, next rune) (code string, err error) {

	current := string(r)
	prevSymbol := string(prev)
	// nextSymbol := string(next)

	if current == "\\" && prevSymbol == "\\" {
		code = "default"

		return code, nil
	}

	if current == "\\" {
		code = "backslash"

		return code, nil
	}

	if unicode.IsDigit(r) {

		if prevSymbol == "\\" {
			code = "default"

			return code, nil
		}

		if unicode.IsDigit(next) {
			return code, errors.New("Wrong string")
		}

		if prev == 0 {
			return code, errors.New("Wrong string")
		}

		return "num", nil
	}

	if prevSymbol == "\\" {
		return code, errors.New("Wrong string")
	}

	return code, nil
}

func processDefault(r rune, b *strings.Builder) {
	b.WriteString(string(r))
}

func processNum(r, prev rune, b *strings.Builder) {
	n, _ := strconv.Atoi(string(r))

	if n == 0 {

		current := b.String()
		trimmed := strings.TrimSuffix(current, string(prev))

		b.Reset()
		b.WriteString(trimmed)

		return
	}

	b.WriteString(repeat(string(prev), n))
}

func unpack(s string) string {
	if s == "" {
		return s
	}

	b := strings.Builder{}

	for i, r := range s {

		prevIdx := i - 1
		nextIdx := i + 1

		var prev, next rune

		if prevIdx >= 0 {
			prev = rune(s[prevIdx])
		}

		if nextIdx <= len(s)-1 {
			next = rune(s[nextIdx])
		}

		t, err := checkSymbol(r, prev, next)

		if err != nil {
			return ""
		}

		switch t {
		case "default":
			processDefault(r, &b)

		case "backslash":
			continue

		case "num":
			processNum(r, prev, &b)
		default:
			processDefault(r, &b)
		}

	}

	return b.String()
}
