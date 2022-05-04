package unpackstring

import (
	"log"
	"strings"
)

func Unpack(s string) string {
	if s == "" {
		return s
	}

	b := strings.Builder{}

	for i, r := range s {

		var prev rune = savePickByIndex(&s, i-1)

		t, err := checkSymbol(i, &s)

		if err != nil {
			log.Println(err.Error())

			return ""
		}

		switch t {
		case CodeDefault:
			processDefault(r, &b)

		case CodeBackslash:
			continue

		case CodeNumber:
			processNum(r, prev, &b)
		default:
			processDefault(r, &b)
		}

	}

	return b.String()
}
