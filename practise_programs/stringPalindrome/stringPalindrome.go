package main

import "fmt"

func checkPalindrome(str string) bool {
	runes := []rune(str)

	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		if runes[i] != runes[j] {
			return false
		}
	}
	return true
}

func main() {
	str := "ananah"

	if checkPalindrome(str) {
		fmt.Println("The word is a palindrome")
	} else {
		fmt.Println("This is not a palindrome word")
	}
}
