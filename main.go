package main

import (
	"alGO/sorting"
	"fmt"
)

func main() {
	// test.TestLinkedList()
	// test.TestStack()
	// test.TestQueue()
	arr := []int{2, 1, 4, 7, 3, 6}
	fmt.Println(sorting.BubbleSort(arr))
	fmt.Println(sorting.Merged(arr))
	fmt.Println(sorting.Insertion(arr))
	fmt.Println(sorting.QuickSort(arr))
	fmt.Println(sorting.Selection(arr))
}
