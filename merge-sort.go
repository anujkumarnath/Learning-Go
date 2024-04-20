package main

import "fmt"

func main() {
  arr := []int{35, 45, 56, 67, 87, 12, 98, 35, 45}
  fmt.Println("unsorted :", arr)

  mergeSort(arr, 0, len(arr) - 1)

  fmt.Println("sorted   :", arr)
}

/* TODO: try and learn iterative version as well */
func mergeSort(arr []int, low, high int) {
  if low >= high {
    return
  }

  mid := (low + high)/2
  // to avoid interger overflow => mid := low + (high - low)/2
  mergeSort(arr, low, mid)
  mergeSort(arr, mid + 1, high)
  merge(arr, low, mid, high)
}

func merge(arr []int, low, mid, high int) {
  temp := make([]int, high - low + 1)
  left := low
  right := mid + 1
  i := 0

  /* in other language, we can do temp[i++] = arr[left++/right++]
    and preferrably use a while loop instead of for to make it look clean
  */
  for ; left <= mid && right <= high; i++ {
    if arr[left] <= arr[right] {
      temp[i] = arr[left]
      left++
    } else {
      temp[i] = arr[right]
      right++
    }
  }

  /* in other language, we can do temp[i++] = arr[left++] */
  for ; left <= mid; left++ {
    temp[i] = arr[left]
    i++
  }

  /* in other language, we can do temp[i++] = arr[right++] */
  for ; right <= high; right++ {
    temp[i] = arr[right]
    i++
  }

  for i := low; i <= high; i++ {
    arr[i] = temp[i - low]
  }
}
