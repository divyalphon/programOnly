package main

import "fmt"

func main() {
	items := []int{12, 56, 12, 69, 69, 35}
	findDuplicatesUsingBruteForce(items)
}
func findDuplicatesUsingBruteForce(items []int) {
	for i := 0; i < len(items); i++ {
		for j := i + 1; j < len(items); j++ {
			if items[i] == items[j] {
				fmt.Println(items[i])
			}
		}
	}
}
