package unpackstring

import (
	"strconv"
	"strings"
)

func repeat(s string, n int) string {

	if n == 0 {
		return ""
	}

	if len(s) > 1 {
		return ""
	}

	return strings.Repeat(s, n-1)
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
