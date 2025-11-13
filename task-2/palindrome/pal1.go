package palindrome

func PalindromeFirst(s string) bool {

	if s != "" {
		for i := 0; i <= len(s)/2; i++ {
			if s[i] != s[len(s)-1-i] {
				return false
			}
		}
	}
	
	return true
}


func PalindromeSecond(s string) bool {
	s1 := ""
	s2 := ""
	if s != "" {

		for i := 0; i < len(s)/2; i++ {
			s1 += string(s[i])
		}

		if len(s)%2 == 0 {
			for i := len(s) - 1; i >= len(s)/2; i-- {
				s2 += string(s[i])
			}
		} else {
			for i := len(s) - 1; i > len(s)/2; i-- {
				s2 += string(s[i])
			}
		}

		if s1 != s2 {
			return false
		}

	}
	return true
}


func PalindromeThird(s string, left int, right int) bool {
	if left >= right {
		return true
	}

	if s[left] != s[right] {
		return false
	}

	return PalindromeThird(s, left+1, right-1)
}