package unpackstring

func savePickByIndex(str *string, i int) rune {

	if i < 0 || i > len((*str))-1 {
		return 0
	}

	return rune((*str)[i])
}
