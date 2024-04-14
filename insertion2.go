package main

import "fmt"

func main() {
	arr := [10]int{4, 69, 34, 61, 0, 56, 64, 99, 12, 11}
	fmt.Println("unsorted:", arr);

	var j int
	for i := 1; i < len(arr); i++ {
		target := arr[i]
		for j = i - 1; j >= 0; j-- {
			if arr[j] <= target {
				break
			}
			arr[j + 1] = arr[j]
		}
		arr[j + 1] = target
	}

	fmt.Println("sorted:", arr);
}
