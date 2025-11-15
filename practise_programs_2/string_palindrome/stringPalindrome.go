package main

import "fmt"

func checkPalindrome(str string) bool {
	runes := []rune(str)
	for i, j := 0, len(str)-1; i < j; i, j = i+1, j-1 {
		if runes[i] != runes[j] {
			return false
		}
	}
	return true
}

func main() {
	str := "hamah"
	if checkPalindrome(str) {
		fmt.Println("the mentioned string is a palindrome")
	} else {
		fmt.Println("the mentioned string is not a palindrome string")
	}
}
