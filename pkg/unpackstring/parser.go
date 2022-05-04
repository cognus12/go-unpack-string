package unpackstring

import (
	"errors"
	"unicode"
)

const BackslashRune = 92
const BackslashChar = "\\"

const CodeDefault = "default"
const CodeBackslash = "backslash"
const CodeNumber = "num"

func isEscaped(str *string, i int) bool {
	prev := savePickByIndex(str, i-1)

	if prev == BackslashRune {
		return true
	}

	return false
}

func checkSymbol(idx int, s *string) (code string, err error) {
	current := savePickByIndex(s, idx)
	prev := savePickByIndex(s, idx-1)
	next := savePickByIndex(s, idx+1)

	currentSymbol := string(current)
	prevSymbol := string(prev)

	if currentSymbol == "\\" && prevSymbol == "\\" {
		code = CodeDefault

		return code, nil
	}

	if currentSymbol == "\\" && !isEscaped(s, idx) {
		code = CodeBackslash

		return code, nil
	}

	if unicode.IsDigit(current) {

		if prevSymbol == "\\" && !isEscaped(s, idx-1) {
			code = CodeDefault

			return code, nil
		}

		if unicode.IsDigit(next) {
			return code, errors.New("Wrong string")
		}

		if prev == 0 {
			return code, errors.New("Wrong string")
		}

		return CodeNumber, nil
	}

	if prevSymbol == BackslashChar {
		return code, errors.New("Wrong string")
	}

	return code, nil
}
