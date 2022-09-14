package selectSort

import "github.com/raymond/sort-golang/bubbleSort"

func SelectSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}
	for i := 0; i < len(arr); i++ {
		var minIndex int = i
		for j := i + 1; j < len(arr); j++ {
			if arr[j] < arr[minIndex] {
				minIndex = j
			}
		}
		arr = bubbleSort.Swap(arr, i, minIndex)
	}
	return arr
}
