package main

import "fmt"

func checkAnagram(str1, str2 string) bool {
	if len(str1) != len(str2) {
		return false
	}
	runeMap := make(map[rune]int)

	for _, ch := range str1 {
		runeMap[ch]++
	}

	for _, ch := range str2 {
		runeMap[ch]--
		if runeMap[ch] < 0 {
			return false
		}
	}

	for _, v := range runeMap {
		if v != 0 {
			return false
		}
	}
	return true
}

func main() {
	str1 := "sachinjangamy"
	str2 := "jangamsacihn"

	if checkAnagram(str1, str2) {
		fmt.Println("strings are anagram")
	} else {
		fmt.Println("strings are not anagram")
	}
}
