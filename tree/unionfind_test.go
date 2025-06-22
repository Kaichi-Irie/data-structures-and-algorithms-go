package tree

import (
	"math/rand"
	"testing"
)

func TestUnionFind(t *testing.T) {
	size := 10_000
	uf := NewUnionFind(size)

	for i := range size {
		for j := i + 1; j < size; j++ {
			if i%10 == j%10 {
				uf.unite(i, j)
			}
		}
	}
	for i := range size {
		for j := i + 1; j < size; j++ {
			if i%10 == j%10 {
				if !uf.same(i, j) {
					t.Errorf("UnionFind test failed: %d and %d should be connected", i, j)
				}
			} else {
				if uf.same(i, j) {
					t.Errorf("UnionFind test failed: %d and %d should not be connected", i, j)
				}
			}
		}
	}
}

func BenchmarkUnionFind(b *testing.B) {
	n_sample := 10_000_000
	size := 100
	uf := NewUnionFind(size)

	// connect random pairs
	for range n_sample {
		i := rand.Intn(size)
		j := rand.Intn(size)
		uf.unite(i, j)
	}
	for range n_sample {
		j := rand.Intn(size)
		uf.same(0, j)
	}
}
