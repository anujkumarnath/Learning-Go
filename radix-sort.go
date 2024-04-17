package main

import "fmt"

func main() {
	/* TODO: find a way to make arr const */
	arr := []int{344, 764, 3, 5, 1234, 9, 8, 223, 2077, 100, 7592384759857, 111343, 144}

	fmt.Println("usorted: ", arr)

	/*
	 * Steps:
	 * 1. Find the maximum number of digits(d) used by the numbers in the array
	 * 2. Perform digit-wise(lsd to msd) counting sort - d iterations
	 * The output of the final iteration is the sorted array of numbers 
	 */

	max := arr[0]
	for i := 1; i < len(arr); i++ {
		if arr[i] > max {
			max = arr[i]
		}
	}

	/* perform counting sort max-no-of-digit times */
	/* #pass = #digts in max = max power of 10 we can devide max to get a non-zero value */
	for placeValue := 1; max/placeValue > 0; placeValue *= 10 {
	   arr = countingSort(arr, placeValue);
	}

	fmt.Println("sorted:  ", arr)
}

/* placeValue means 1, 10, 100,... */
func countingSort(arr []int, placeValue int) []int {
	count := make([]int, 10);

	for i := len(arr) - 1; i >= 0; i-- {
		/* 7[1]2 -> ([712/(10*10)]%10) -> 1 */
		count[arr[i]/placeValue % 10]++
	}

	for i := 1; i < len(count); i++ {
		count[i] += count[i - 1]
	}

	b := make([]int, len(arr))

	for i := len(arr) - 1; i >= 0; i-- {
		digit := arr[i]/placeValue % 10
		count[digit]--
		b[count[digit]] = arr[i]
	}

	return b
}
