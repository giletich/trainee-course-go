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
    runes := []rune(s)
    
    for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
        runes[i], runes[j] = runes[j], runes[i]
    }
    
    return string(runes)
}