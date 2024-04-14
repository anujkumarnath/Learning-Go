package main

import "fmt"

func main() {
	arr := [10]int{45, 13, 45, 11, 51, 5, 16, 61, 84, 1}
	fmt.Println("unsorted:", arr)

	for i := 0; i < len(arr); i++ {
		min_index := i
		for j := i + 1; j < len(arr); j++ {
			if arr[j] < arr[min_index] {
				min_index = j
			}
		}
		temp := arr[min_index]
		arr[min_index] = arr[i]
		arr[i] = temp
	}

	fmt.Println("sorted:  ", arr)
}
