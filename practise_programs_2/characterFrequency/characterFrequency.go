package main

import "fmt"

func countCharFreq(str string) map[rune]int {
	runeMap := make(map[rune]int)

	for _, c := range str {
		runeMap[c]++
	}
	return runeMap
}

func main() {
	str := "sachinjangam"
	countChFreq := countCharFreq(str)

	for ch, count := range countChFreq {
		fmt.Printf("%c - %d", ch, count)
		fmt.Println()
	}
}
