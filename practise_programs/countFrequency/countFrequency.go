package main

import (
	"fmt"
	"sort"
)

func countFreq(word string) map[rune]int {
	wordMap := make(map[rune]int)
	//1st way
	for _, ch := range word {
		wordMap[ch]++
	}

	//2nd way
	//runes := []rune(word)
	//for i := 0; i < len(runes); i++ {
	//	wordMap[runes[i]]++
	//}
	return wordMap
}

func sortSlice(word string) string {
	keys := []rune(word)

	sort.Slice(keys, func(i, j int) bool {
		return keys[i] < keys[j]
	})
	return string(keys)
}

func main() {
	str := "sachinjangam"
	word := countFreq(str)
	fmt.Println("Count with the characters: (non sorted order)")
	for key, value := range word {
		fmt.Printf("%c : %d", key, value)
		fmt.Println()
	}

	sorted := sortSlice(str)
	fmt.Println("Keys in sorted order")
	printed := make(map[rune]bool)
	for _, ch := range sorted {
		if !printed[ch] {
			fmt.Printf("%c: %d", ch, word[ch])
			printed[ch] = true
			fmt.Println()
		}
	}
}
