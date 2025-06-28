package tree

type Heap []int

func BuildMinHeap(array []int) *Heap {
	h := Heap(array)
	length := len(h)
	for idx := length/2 + 1; idx > 0; idx-- {
		h.Heapify(idx, length)
	}
	return &h
}

// 1-indexed min heap
func Left(i int) int {
	return 2 * i
}
func Right(i int) int {
	return 2*i + 1
}

func Exchange(array []int, i int, j int) {
	tmp_value := array[i]
	array[i] = array[j]
	array[j] = tmp_value
}

func (h *Heap) Heapify(idx int, length int) {
	if length == 0 {
		return
	}
	value := (*h)[idx-1]
	leftIndex := Left(idx)
	rightIndex := Right(idx)

	var leftValue int
	if leftIndex <= length {
		leftValue = (*h)[leftIndex-1]
	} else {
		leftValue = value
	}

	var rightValue int
	if rightIndex <= length {
		rightValue = (*h)[rightIndex-1]
	} else {
		rightValue = value
	}

	IndexOfSmallest := idx
	if leftValue < value {
		IndexOfSmallest = leftIndex
	}
	if rightValue < min(value, leftValue) {
		IndexOfSmallest = rightIndex
	}

	if IndexOfSmallest != idx {
		Exchange(*h, idx-1, IndexOfSmallest-1)
		(*h).Heapify(IndexOfSmallest, length)
	}
}

func (h *Heap) Pop() int {
	length := len(*h)
	if length == 0 {
		return 0
	}
	minValue := (*h)[0]
	if length == 1 {
		(*h) = (*h)[:0] // empty the heap
		return minValue
	}
	Exchange(*h, 0, length-1)
	(*h).Heapify(1, length-1)
	(*h) = (*h)[:length-1] // remove last element
	return minValue
}
