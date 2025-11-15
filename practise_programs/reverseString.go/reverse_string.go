package main

import "fmt"

func reverse(name string) string {
	runes := []rune(name)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func main() {

	// simple logic

	//name := "sachin"
	//runes := []rune(name)
	//for i := len(runes) - 1; i > 0; i-- {
	//	fmt.Printf("%c", runes[i])
	//}

	//Two pointer algorithm
	name := "sachin"
	reversedString := reverse(name)
	fmt.Println("the reversed string is ", reversedString) // if you want to print string value

	// if you want to print rune in bytes.

	fmt.Printf("the bytes %c", []rune(reversedString))
}
