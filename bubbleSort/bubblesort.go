package bubbleSort

func BubbleSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	for e := len(arr) - 1; e > 0; e-- {
		for i := 0; i < e; i++ {
			if arr[i] > arr[i+1] {
				Swap(arr, i, i+1)
			}
		}
	}

	return arr
}

func Swap(arr []int, i, j int) []int {
	temp := arr[j]
	arr[j] = arr[i]
	arr[i] = temp
	return arr
}
