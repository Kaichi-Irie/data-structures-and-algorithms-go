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
