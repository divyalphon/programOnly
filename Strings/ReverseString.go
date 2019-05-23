package main

import "fmt"

func main() {
	strone := "Divya"
	fmt.Println(reverse(strone))
}
func reverse(name string) string {
	runes := []rune(name)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
