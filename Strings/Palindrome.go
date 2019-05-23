package main

import (
	"fmt"
	"strings"
)

func main() {
	strone := "nitin"
	//strone := "Divya"
	newvalue := reverse(strone)
	if strings.EqualFold(strone, newvalue) {
		fmt.Println("Strings are Palindrome")
	} else {
		fmt.Println("Strings are  NOT Palindrome")
	}
}

func reverse(strone string) string {
	runes := []rune(strone)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
