package main

import "fmt"

func findNonRepeatChar(str string) rune {
	runeMap := make(map[rune]int)

	for _, ch := range str {
		runeMap[ch]++
	}

	for _, ch := range str {
		if runeMap[ch] == 1 {
			return ch
		}
	}
	return 0
}

func main() {
	str := "sachinjangam"
	ch := findNonRepeatChar(str)
	fmt.Printf("\n first non repeating character is %c ", ch)
}
