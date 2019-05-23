package main

import "fmt"

func main() {
	items := []int{12, 23, 69, 89, 69, 1}
	findDuplicatesHashMap(items)
}

func findDuplicatesHashMap(items []int) {
	mapone := make(map[int]int)
	for _, a := range items {
		if mapone[a] == 0 {
			mapone[a] = 1
		} else {
			mapone[a] = mapone[a] + 1
		}
	}
	for key, value := range mapone {
		if value > 1 {
			fmt.Println(key, " =", mapone[key])
		}
	}

}
