package binary

import (
	"fmt"
	"math/rand"
	"testing"
)

func TestIsPeak(t *testing.T) {
	tests := []struct {
		a    []int
		i    int
		want bool
	}{
		{[]int{1, 2, 3, 4, 5}, 0, false},
		{[]int{1, 2, 3, 4, 5}, 1, false},
		{[]int{1, 2, 3, 4, 5}, 2, false},
		{[]int{1, 2, 3, 4, 5}, 3, false},
		{[]int{1, 2, 3, 4, 5}, 4, true},
		{[]int{3, 2, 1, 4, 2}, 0, true},
		{[]int{3, 2, 1, 4, 2}, 1, false},
		{[]int{3, 2, 1, 4, 2}, 2, false},
		{[]int{3, 2, 1, 4, 2}, 3, true},
		{[]int{3, 2, 1, 4, 2}, 4, false},
	}
	for _, test := range tests {
		if got := IsPeak(test.a, test.i); got != test.want {
			t.Errorf("IsPeak(%v,%d) = %v; want %v", test.a, test.i, got, test.want)
		}
	}
}

func TestFindPeak(t *testing.T) {
	tests := []struct {
		a    []int
		want int
	}{
		{[]int{1, 2, 3, 4, 5}, 4},
		{[]int{5, 4, 3, 2, 1}, 0},
		{[]int{1, 2, 3, 4, 5, 4}, 4},
	}
	for _, test := range tests {
		if got := FindPeak(test.a, 0, len(test.a)); got == -1 {
			t.Errorf("FindPeak(%v,0,%d) = %d; want a peak", test.a, len(test.a), got)
		} else if got != test.want {
			t.Errorf("FindPeak(%v,0,%d) = %d; want %d", test.a, len(test.a), got, test.want)
		}
	}

	// generate random int slice and test FindPeak
	sampleSize := 100
	arraySize := 100
	seed := rand.NewSource(1)
	rand := rand.New(seed)
	a := make([]int, arraySize)
	for range sampleSize {
		for i := range a {
			a[i] = rand.Intn(100)
		}
		fmt.Println("a:", a)
		if got := FindPeak(a, 0, len(a)); got == -1 {
			t.Errorf("FindPeak(%v,0,%d) = %d; want a peak", a, len(a), got)
		} else if !IsPeak(a, got) {
			t.Errorf("FindPeak(%v,0,%d) = %d; want a peak", a, len(a), got)
		}
	}
}
