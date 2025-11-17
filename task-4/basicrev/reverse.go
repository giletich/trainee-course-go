package basicreverse

func BasicReverse(s string) string {
	strs := StringDivision(s)
	res := Reverse(strs)
	return res
}

func StringDivision(s string) []string {
	words := []string{}
	word := ""
	for i, value := range s {
		if s[i] == ' ' {
			words = append(words, word)
			word = ""
		} else {
			word += string(value)
		}
	}
	words = append(words, word)
	return words
}

func Reverse(s []string) string {
	result := ""

	for i, value := range s {
		word := ""
		for j := len(value) - 1; j >= 0; j-- {
			word += string(value[j])
		}
		result += word
		if i != len(s)-1 {
			result += " "
		}
	}

	return result
}
