package main

import "fmt"

func main() {
	arr := [10]int{45, 2, 54, 25, 16, 63, 33, 5, 26, 2}
	fmt.Println("unsorted:", arr)

	var i, j int
	for i = 0; i < len(arr); i++ {
		min_index := i
		for j = i + 1; j < len(arr); j++ {
			if arr[j] >= arr[min_index] {
				continue
			}
			min_index = j
		}
		temp := arr[i]
		arr[i] = arr[min_index]
		arr[min_index] = temp
	}
	
	fmt.Println("sorted:", arr)
}
