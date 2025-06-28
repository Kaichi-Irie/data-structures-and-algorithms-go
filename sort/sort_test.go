package sort

import (
	"math/rand"
	"reflect"
	"sort"
	"testing"
)

func CreateTestCases() []struct {
	input    []int
	expected []int
} {
	sortTests := []struct {
		input    []int
		expected []int
	}{
		{[]int{5, 2, 9, 1, 5, 6}, []int{1, 2, 5, 5, 6, 9}},
		{[]int{3}, []int{3}},
		{[]int{}, []int{}},
		{[]int{1, 2, 3}, []int{1, 2, 3}},
		{[]int{3, 2, 1}, []int{1, 2, 3}},
	}
	// generate random test cases
	var arr, sortedArr []int
	for i := 0; i < 200; i++ {
		arr = generateRandomSlice(10000)
		sortedArr = make([]int, len(arr))
		copy(sortedArr, arr)
		sort.Ints(sortedArr)
		sortTests = append(sortTests, struct {
			input    []int
			expected []int
		}{arr, sortedArr})
	}
	return sortTests
}

func generateRandomSlice(n int) []int {
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		arr[i] = rand.Intn(1000) // Random integers between 0 and 999
	}
	return arr
}
func TestInsertionSort(t *testing.T) {
	for _, test := range CreateTestCases() {

		t.Run("Insertion Sort", func(t *testing.T) {
			InsertionSort(test.input)
			if !reflect.DeepEqual(test.input, test.expected) {
				t.Errorf("Expected %v, got %v", test.expected, test.input)
			}
		})
	}
}

func TestMerge(t *testing.T) {
	tests := []struct {
		input    []int
		p        int
		q        int
		r        int
		expected []int
	}{
		{[]int{1, 4, 6, 7, 2, 3, 8, 9, 5}, 0, 4, 8, []int{1, 2, 3, 4, 6, 7, 8, 9, 5}}}
	for _, test := range tests {
		t.Run("Merge", func(t *testing.T) {
			Merge(test.input, test.p, test.q, test.r)
			if !reflect.DeepEqual(test.input, test.expected) {
				t.Errorf("Expected %v, got %v", test.expected, test.input)
			}
		})
	}

}
func TestMergeSort(t *testing.T) {
	for _, test := range CreateTestCases() {
		t.Run("MergeSort", func(t *testing.T) {
			MergeSort(test.input, 0, len(test.input))
			if !reflect.DeepEqual(test.input, test.expected) {
				t.Errorf("Expected %v, got %v", test.expected, test.input)
			}
		})
	}
}

func TestPartition(t *testing.T) {
	tests := []struct {
		input    []int
		l        int
		r        int
		expected int
	}{
		{[]int{5, 2, 9, 1, 5, 6}, 0, 6, 4},
		{[]int{3, 2, 1}, 0, 3, 0},
		{[]int{1}, 0, 1, 0},
		{[]int{}, 0, 0, -1},
	}

	for _, test := range tests {
		t.Run("Partition", func(t *testing.T) {
			result := Partition(test.input, test.l, test.r)
			if result != test.expected {
				t.Errorf("Expected %d, got %d", test.expected, result)
			}
		})
	}

}

func TestDeterministicQuickSort(t *testing.T) {
	// deterministic quicksort
	for _, test := range CreateTestCases() {
		t.Run("QuickSort", func(t *testing.T) {
			QuickSort(test.input, 0, len(test.input), true)
			if !reflect.DeepEqual(test.input, test.expected) {
				t.Errorf("Expected %v, got %v", test.expected, test.input)
			}
		})
	}
}

func TestRandomizedQuickSort(t *testing.T) {
	// randomized quicksort
	for _, test := range CreateTestCases() {
		t.Run("QuickSort", func(t *testing.T) {
			QuickSort(test.input, 0, len(test.input), false)
			if !reflect.DeepEqual(test.input, test.expected) {
				t.Errorf("Expected %v, got %v", test.expected, test.input)
			}
		})
	}
}

func TestHeapSort(t *testing.T) {
	for _, test := range CreateTestCases() {
		t.Run("HeapSort", func(t *testing.T) {
			HeapSort(test.input)
			if !reflect.DeepEqual(test.input, test.expected) {
				t.Errorf("Expected %v, got %v", test.expected, test.input)
			}
		})
	}
}
