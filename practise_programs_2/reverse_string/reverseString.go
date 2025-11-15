package main

import "fmt"

func reverseString(str string) []rune {
	runes := []rune(str)

	for i, j := 0, len(str)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return runes
}

func main() {
	str := "my name is sachin jangam"
	reversed := string(reverseString(str))
	fmt.Printf("the reversed string is %s", reversed)

}
