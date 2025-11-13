package main

import (
	"fmt"

	"task-2/palindrome"
)

func main() {
	s := "oooo"
	fmt.Println(palindrome.PalindromeFirst(s))
	fmt.Println(palindrome.PalindromeSecond(s))
	fmt.Println(palindrome.PalindromeThird(s, 0, len(s)-1))
}
