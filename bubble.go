package main

import "fmt"

func main() {
	arr := [10]int{3, 45, 1, 67, 56, 88, 99, 45, 28, 4}

	fmt.Println("unsorted:", arr)

	for i := 1; i < len(arr); i++ {
		for j := 0; j < len(arr) - i; j++ {
			if arr[j] <= arr[j + 1] {
				continue
			} else {
				temp := arr[j]
				arr[j] = arr[j + 1]
				arr[j + 1] = temp
			}
		}
	}

	fmt.Println("sorted:", arr)
}
