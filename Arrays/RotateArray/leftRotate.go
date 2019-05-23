package main

import "fmt"

func main() {
	items := []int{12, 36, 89, 78, 45, 23}
	fmt.Println("Before Left Rotate ", items)
	leftrotate(items, 2)

	fmt.Println("After Left Rotate ", items)
}
func leftrotate(items []int, postion int) {
	var temp int

	for i := 0; i < postion; i++ {
		temp = items[0]
		for j := 0; j < len(items)-1; j++ {
			items[j] = items[j+1]
		}
		items[len(items)-1] = temp
	}
}
