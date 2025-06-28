package sort

import (
	"data-structures-and-algorithms-go/tree"
)

func HeapSort(array []int) {
	if len(array) == 0 {
		return
	}
	heap := tree.BuildMinHeap(array)
	length := len(*heap)
	for length > 0 {
		// set minimum value to the last
		tree.Exchange(*heap, 0, length-1)
		length--
		heap.Heapify(1, length)
	}

	// reverse the array to get sorted order
	for i := 0; i < len(array)/2; i++ {
		tree.Exchange(array, i, len(array)-1-i)
	}
}
