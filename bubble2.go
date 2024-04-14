package main

import "fmt"

func main() {
	arr := [10]int{34, 54, 23, 1, 55, 77, 25, 9, 90, 23}
	fmt.Println("unsorted:", arr)

	for i := 1; i < len(arr); i++ {
		for j := 0; j < len(arr) - i; j++ {
			if arr[j] <= arr[j + 1] {
				continue;
			}
			temp := arr[j]
			arr[j] = arr[j + 1]
			arr[j + 1] = temp;
		}
	}

	fmt.Println("sorted:", arr)
}
