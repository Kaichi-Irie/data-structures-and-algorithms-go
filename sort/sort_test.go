package sort

import (
	"reflect"
	"testing"
)

var sortTests = []struct {
	input    []int
	expected []int
}{
	{[]int{5, 2, 9, 1, 5, 6}, []int{1, 2, 5, 5, 6, 9}},
	{[]int{3}, []int{3}},
	{[]int{}, []int{}},
	{[]int{1, 2, 3}, []int{1, 2, 3}},
	{[]int{3, 2, 1}, []int{1, 2, 3}}}

func TestInsertionSort(t *testing.T) {
	for _, test := range sortTests {

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
	for _, test := range sortTests {
		t.Run("MergeSort", func(t *testing.T) {
			MergeSort(test.input, 0, len(test.input))
			if !reflect.DeepEqual(test.input, test.expected) {
				t.Errorf("Expected %v, got %v", test.expected, test.input)
			}
		})
	}
}
