package main

import (
	"fmt"
	"sort"
	"strings"
)

func countFreWord(wd string) map[string]int {
	wordMap := make(map[string]int)

	str := strings.ToLower(wd)
	words := strings.Fields(str)

	for _, word := range words {
		wordMap[word]++
	}
	return wordMap
}

func sortedString(wd map[string]int) []string {
	keys := make([]string, 0, len(wd))

	for key := range wd {
		keys = append(keys, key)
	}

	sort.Strings(keys)
	return keys
}

func main() {
	str := "I love go and i love programming in go"
	wordsCount := countFreWord(str)
	for key, value := range wordsCount {
		fmt.Printf(" %s : %d", key, value)
		fmt.Println()
	}

	fmt.Println("In sorted order")
	sorted := sortedString(wordsCount)
	fmt.Println(sorted)
}
