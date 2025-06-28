package tree

import (
	"testing"
)

func TestHeap(t *testing.T) {
	// Test cases for the heap implementation
	tests := []struct {
		name    string
		input   []int
		minimum int
	}{
		// {"Empty Heap", []int{}, 0},
		{"Single Element", []int{5}, 5},
		{"Three Elements", []int{3, 1, 5}, 1},
		{"Multiple Elements", []int{4, 10, 3, 5, 1}, 1},
		{"Many Elements", []int{7, 2, 8, 1, 6, 3, 5, 4, 9}, 1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := BuildMinHeap(tt.input)
			if got := (*h)[0]; got != tt.minimum {
				t.Errorf("Heap minimum = %v, want %v", got, tt.minimum)
			}
		})
	}
}

func TestHeapPop(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected []int
	}{
		{"Single Element", []int{5}, []int{5}},
		{"Two Elements", []int{3, 1}, []int{1, 3}},
		{"Multiple Elements", []int{4, 10, 3, 5, 1}, []int{1, 3, 4, 5, 10}},
		{"Many Elements", []int{7, 2, 8, 1, 6, 3, 5, 4, 9}, []int{1, 2, 3, 4, 5, 6, 7, 8, 9}},
		{"Already Sorted", []int{1, 2, 3, 4, 5}, []int{1, 2, 3, 4, 5}},
		{"Reverse Sorted", []int{5, 4, 3, 2, 1}, []int{1, 2, 3, 4, 5}},
		{"Duplicates", []int{3, 1, 2, 3, 1, 2}, []int{1, 1, 2, 2, 3, 3}},
		{"Negative Numbers", []int{-1, -3, -2, 0, 2, 1}, []int{-3, -2, -1, 0, 1, 2}},
		{"Mixed Numbers", []int{3, -1, 2, 5, -3, 4}, []int{-3, -1, 2, 3, 4, 5}},
		{"Empty Array", []int{}, []int{}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := BuildMinHeap(tt.input)
			for i := 0; i < len(tt.input); i++ {
				minValue := h.Pop()
				if minValue != tt.expected[i] {
					t.Errorf("HeapPop() = %v, want %v", minValue, tt.expected[i])
				}
			}
		})
	}
}
