package main

import (
	"fmt"

	"github.com/raymond/sort-golang/bubbleSort"
	"github.com/raymond/sort-golang/selectSort"
)

func main() {
	arr := []int{1, 3, 3, 2, 14, 4}

	fmt.Println(bubbleSort.BubbleSort(arr))
	fmt.Println(selectSort.SelectSort(arr))
}
