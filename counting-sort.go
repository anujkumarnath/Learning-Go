package main

import "fmt"

func main() {
	/* TODO: find a way to make arr const */
	arr := [10]int{4, 4, 3, 5, 1, 9, 8, 3, 2, 1}

	fmt.Println("usorted: ", arr)

	/* should replace with -ve most int
	 * and start with index 0 to avoid accessing arr[0]
	 * if doesn't exist
	 */
	max := arr[0]
	for i := 1; i < len(arr); i++ {
		if arr[i] > max {
			max = arr[i]
		}
	}

	/* TODO: learn slices */
	counts := make([]int, max + 1)
	for i := 0; i < len(arr); i++ {
		counts[arr[i]]++
	}

	for i := 1; i < len(counts); i++ {
		counts[i] += counts[i - 1]
	}

	var b [len(arr)]int

	for i := len(arr) - 1; i >= 0; i-- {
		/* go doesn't support b[--counts[arr[i]]] = arr[i]
		 * so splitting into two lines
		 */
		counts[arr[i]]--
		b[counts[arr[i]]] = arr[i]
	}

	for i := 0; i < len(arr); i++ {
		arr[i] = b[i]
	}

	fmt.Println("sorted:  ", arr)
}
