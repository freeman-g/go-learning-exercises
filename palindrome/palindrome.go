package main

import (
	"strings"
)

func isPalindrome(word string) bool {
	letters := strings.Split(word, "")
	length := len(letters)
	half := length / 2

	for i := 0; i <= half; i++ {
		if letters[i] != letters[length-(i+1)] {
			return false
		}
	}
	return true
}
