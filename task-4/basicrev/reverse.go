package basicreverse

func BasicReverse(s string) string {
	strs := StringDivision(s)
	res := Reverse(strs)
	return res
}

func StringDivision(s string) []string {
    words := []string{}
    word := ""
    runes := []rune(s)
    
    for i := 0; i < len(runes); i++ {
        if runes[i] == ' ' {
            if word != "" {
                words = append(words, word)
                word = ""
            }
        } else {
            word += string(runes[i])
        }
    }
    
    if word != "" {
        words = append(words, word)
    }
    
    return words
}

func Reverse(s []string) string {
    result := ""
    
    for i, value := range s {
        runes := []rune(value)
        word := ""
        
        for j := len(runes) - 1; j >= 0; j-- {
            word += string(runes[j])
        }
        
        result += word
        
        if i != len(s)-1 {
            result += " "
        }
    }
    
    return result
}
