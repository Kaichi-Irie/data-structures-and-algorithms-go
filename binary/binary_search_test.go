package binary

import "testing"

func TestBinarySearch(t *testing.T) {
	tests := []struct {
		A        []int
		n1       int
		n2       int
		k        int
		expected bool
	}{
		{[]int{4, 2, 3, 1, 5}, 0, 5, 3, true},
		{[]int{4, 2, 3, 1, 5}, 0, 5, 0, false},
		{[]int{4, 2, 3, 1, 5}, 0, 5, 6, false},
		{[]int{4, 2, 3, 1, 5}, 0, 5, -1, false},
		{[]int{9, 3, 5, 7, 1, 2, 4, 6, 8}, 0, 9, 3, true},
	}

	for _, test := range tests {
		result := BinarySearch(test.A, test.n1, test.n2, test.k)
		if result != test.expected {
			t.Errorf("BinarySearch(%v, %d, %d, %d) = %v; expected %v", test.A, test.n1, test.n2, test.k, result, test.expected)
		}
	}
}
