package main

import "fmt"

func main() {
	items := []int{12, 36, 89, 78, 24}
	fmt.Println("Before Right Rotate ", items)
	rightrotate(items, 3)
	fmt.Println("After Right Rotate ", items)
}

func rightrotate(items []int, position int) {
	var temp int

	for i := 1; i < position; i++ {
		temp = items[len(items)-1]
		for j := len(items) - 1; j > 0; j-- {
			items[j] = items[j-1]
		}
		items[0] = temp
	}
}
