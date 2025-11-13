package smartreverse

import "strings"

func SmartReverse(s string) string {
	strs := strings.Split(s, " ")
	words := []string{}

	for _, value := range strs {
		words = append(words, Reverse(value))
	}

	res := strings.Join(words, " ")
	return res
}

func Reverse(s string) string {

		word := ""

		for i := len(s) - 1; i >= 0; i-- {
			word += string(s[i])
		}
		
		return word
	}
