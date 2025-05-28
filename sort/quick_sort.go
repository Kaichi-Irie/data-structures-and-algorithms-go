package sort

import (
	"math/rand"
)

/*
worst time complexity: O(n^2)
average time complexity: O(n log n)
space complexity: O(1) (in-place sorting) (not counting recursion stack space)
*/

func Partition(arr []int, l int, r int) int {
	if r-l <= 0 {
		return -1
	} else if r-l == 1 {
		return l
	}
	x := arr[r-1]
	i := l
	for j := l; j < r-1; j++ {
		if arr[j] <= x {
			Exchange(arr, i, j)
			i++
		}
	}
	Exchange(arr, i, r-1)
	return i
}

func randRange(min, max int) int {
	val := rand.Intn(max-min) + min
	return val
}

func RandomPartition(arr []int, l int, r int) int {
	if r-l <= 0 {
		return -1
	} else if r-l == 1 {
		return l
	}
	i := randRange(l, r)
	Exchange(arr, i, r-1)
	return Partition(arr, l, r)
}

func Exchange(arr []int, i int, j int) {
	if min(i, j) < 0 || max(i, j) >= len(arr) {
		panic("index out of range")
	}
	tmp := arr[i]
	arr[i] = arr[j]
	arr[j] = tmp
}

func QuickSort(arr []int, l int, r int, deterministic bool) {
	if r-l <= 1 {
		return
	}
	var m int
	if deterministic {
		m = Partition(arr, l, r)
	} else {
		m = RandomPartition(arr, l, r)
	}
	if m == -1 {
		panic("partition failed")
	}
	QuickSort(arr, l, m, deterministic)
	QuickSort(arr, m+1, r, deterministic)
}
