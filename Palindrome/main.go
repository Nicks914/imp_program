package main

import (
	"fmt"
	"strings"
)

func main() {
	test := "Madam"
	if isPalindrome(test) {
		fmt.Println("Is a Palindrome")

	} else {
		fmt.Println("Is not a Palindrome")
	}

	if Palindrome(test) {
		fmt.Println("String is a Palindrome")

	} else {
		fmt.Println("String not a Palindrome")
	}
}

func isPalindrome(s string) bool {

	s = strings.ToLower(s)
	s = strings.ReplaceAll(s, " ", "")

	runes := []rune(s)
	length := len(runes)

	for i := 0; i <= length/2; i++ {
		if runes[i] != runes[length-1-i] {
			return false
		}
	}
	return true
}

func Palindrome(s string) bool {

	s = strings.ToLower(s)
	runes := []rune(s)

	left, right := 0, len(runes)-1

	for left < right {
		if runes[left] != runes[right] {
			return false
		}
		left++
		right--
	}
	return true

}
