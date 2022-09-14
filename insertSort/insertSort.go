package insertSort

import "github.com/raymond/sort-golang/bubbleSort"

func InsertSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	for i := 1; i < len(arr); i++ {
		for j := i - 1; j >= 0; j-- {
			if arr[j] > arr[j+1] {
				bubbleSort.Swap(arr, j, j+1)
			}
		}
	}
	return arr
}
