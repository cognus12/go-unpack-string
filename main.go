package main

import (
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	str := `qwe\45`
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

func savePickByIndex(str *string, i int) rune {

	if i < 0 || i > len((*str))-1 {
		return 0
	}

	return rune((*str)[i])
}

func isEscaped(str *string, i int) bool {
	prev := savePickByIndex(str, i-1)

	if prev == 92 {
		return true
	}

	return false
}

func checkSymbolNew(idx int, s *string) (code string, err error) {
	current := savePickByIndex(s, idx)
	prev := savePickByIndex(s, idx-1)
	next := savePickByIndex(s, idx+1)

	currentSymbol := string(current)
	prevSymbol := string(prev)

	if currentSymbol == "\\" && prevSymbol == "\\" {
		code = "default"

		return code, nil
	}

	if currentSymbol == "\\" && !isEscaped(s, idx) {
		code = "backslash"

		return code, nil
	}

	if unicode.IsDigit(current) {

		if prevSymbol == "\\" && !isEscaped(s, idx-1) {
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

// func checkSymbol(r, prev, next rune) (code string, err error) {

// 	current := string(r)
// 	prevSymbol := string(prev)
// 	// nextSymbol := string(next)

// 	if current == "\\" && prevSymbol == "\\" {
// 		code = "default"

// 		return code, nil
// 	}

// 	if current == "\\" {
// 		code = "backslash"

// 		return code, nil
// 	}

// 	if unicode.IsDigit(r) {

// 		if prevSymbol == "\\" {
// 			code = "default"

// 			return code, nil
// 		}

// 		if unicode.IsDigit(next) {
// 			return code, errors.New("Wrong string")
// 		}

// 		if prev == 0 {
// 			return code, errors.New("Wrong string")
// 		}

// 		return "num", nil
// 	}

// 	if prevSymbol == "\\" {
// 		return code, errors.New("Wrong string")
// 	}

// 	return code, nil
// }

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

		var prev rune = savePickByIndex(&s, i-1)

		t, err := checkSymbolNew(i, &s)

		if err != nil {
			log.Println(err.Error())

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
