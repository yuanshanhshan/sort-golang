package quickSort

func QuickSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	middle := arr[0]
	var left []int
	var right []int
	for i := 1; i < len(arr); i++ {
		if arr[i] > middle {
			right = append(right, arr[i])
		} else {
			left = append(left, arr[i])
		}
	}

	middle_s := []int{middle}
	left = QuickSort(left)
	right = QuickSort(right)
	arr = append(append(left, middle_s...), right...)
	return arr
}
