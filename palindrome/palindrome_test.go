package main

import (
	"testing"
)

func TestIsPalindrome(t *testing.T) {

	palindromes := []string{"eye", "abba", "racecar"}

	for _, word := range palindromes {
		test := isPalindrome(word)
		if !test {
			t.Errorf("%s IS a palindrome, function returned %v", word, test)
		}
	}

	nonPalindromes := []string{"kitten", "puppy"}

	for _, word := range nonPalindromes {
		test := isPalindrome(word)
		if test {
			t.Errorf("%s IS NOT a palindrome, function returned %v", word, test)
		}
	}
}
