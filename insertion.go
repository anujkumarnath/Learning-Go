package main

import "fmt"

func main() {
	arr := [10]int{23, 1, 56, 35, 23, 0, 34, 5, 35, 20}
	fmt.Println("unsorted:", arr)

	var i, j int
	for i = 1; i < len(arr); i++ {
		target := arr[i]
		for j = i - 1; j >= 0; j-- {
			if targer >= arr[j] {
				break
			}
			arr[j + 1] = arr[j]
		}
		arr[j + 1] = target
	}

	fmt.Println("sorted:", arr)
}
