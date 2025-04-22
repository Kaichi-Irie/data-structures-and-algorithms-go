package binary

import "testing"

func TestBisectLeft(t *testing.T) {
	tests := []struct {
		A        []int
		k        int
		expected int
	}{
		{[]int{1, 2, 3, 4, 5}, 3, 2},
		{[]int{1, 2, 3, 4, 5}, 0, 0},
		{[]int{1, 2, 3, 4, 5}, 6, 5},
		{[]int{1, 2, 3, 4, 5}, 1, 0},
		{[]int{1, 2, 2, 2, 2, 3, 3, 4, 5}, 2, 1},
		{[]int{1, 2, 2, 2, 2, 3, 3, 4, 5}, 3, 5},
	}

	for _, test := range tests {
		result := BisectLeft(test.A, test.k)
		if result != test.expected {
			t.Errorf("BisectLeft(%v, %d) = %d; expected %d", test.A, test.k, result, test.expected)
		}
	}
}
func TestBisectRight(t *testing.T) {
	tests := []struct {
		A        []int
		k        int
		expected int
	}{
		{[]int{1, 2, 3, 4, 5}, 3, 3},
		{[]int{1, 2, 3, 4, 5}, 0, 0},
		{[]int{1, 2, 3, 4, 5}, 6, 5},
		{[]int{1, 2, 3, 4, 5}, 1, 1},
		{[]int{1, 2, 2, 2, 2, 3, 3, 4, 5}, 2, 5},
		{[]int{1, 2, 2, 2, 2, 3, 3, 4, 5}, 3, 7},
	}

	for _, test := range tests {
		result := BisectRight(test.A, test.k)
		if result != test.expected {
			t.Errorf("BisectRight(%v, %d) = %d; expected %d", test.A, test.k, result, test.expected)
		}
	}
}
