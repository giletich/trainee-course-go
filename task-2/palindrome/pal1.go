package palindrome

func PalindromeFirst(s string) bool {

	runes := []rune(s)

	if string(runes) != "" {

		for i := 0; i <= len(runes)/2; i++ {
			if runes[i] != runes[len(runes)-1-i] {
				return false
			}
		}
	}
	
	return true
}


func PalindromeSecond(s string) bool {
	runes := []rune(s)

	s1 := []rune{}
	s2 := []rune{}
	if s != "" {

		for i := 0; i < len(runes)/2; i++ {
			s1 = append(s1, runes[i]) 
		}

		if len(runes)%2 == 0 {
			for i := len(runes) - 1; i >= len(runes)/2; i-- {
				s2 = append(s2, runes[i]) 
			}
		} else {
			for i := len(runes) - 1; i > len(runes)/2; i-- {
				s2 = append(s2, runes[i]) 
			}
		}

		for i := range s1 {

			if s1[i] != s2[i] {
				return false
			}

		}

	}
	return true
}


func PalindromeThird(s string, left int, right int) bool {

	runes := []rune(s)
	if left >= right {
		return true
	}

	if runes[left] != runes[right] {
		return false
	}

	return PalindromeThird(string(runes), left+1, right-1)
}