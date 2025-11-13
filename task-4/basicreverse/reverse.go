package basicreverse

import "fmt"

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
	fmt.Println(words)
	return words
}

func Reverse(s []string) string {
	result := ""

	for _, value := range s {
		word := ""
		for i := len(value) - 1; i >= 0; i-- {
			word += string(value[i])
		}
		result += word
		result += " "
		fmt.Println(value)
	}

	return result
}