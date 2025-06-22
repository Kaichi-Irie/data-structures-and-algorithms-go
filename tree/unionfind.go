package tree

type UnionFind struct {
	size   int
	parent []int
	rank   []int
}

func (uf *UnionFind) find(x int) int {
	if uf.parent[x] == x {
		return x
	}

	x_root := uf.find(uf.parent[x])
	uf.parent[x] = x_root // path compression
	return x_root
}

func (uf *UnionFind) unite(x int, y int) {
	x_root := uf.find(x)
	y_root := uf.find(y)
	if x_root == y_root {
		return // already connected
	}
	// connect without rank
	// uf.parent[x_root] = uf.parent[y_root]

	// connect with rank
	if uf.rank[x_root] < uf.rank[y_root] {
		uf.parent[x_root] = y_root
		uf.rank[y_root] += 1
	} else {
		uf.parent[y_root] = x_root
		uf.rank[x_root] += 1
	}
}

func (uf *UnionFind) same(x int, y int) bool {
	x_root := uf.find(x)
	y_root := uf.find(y)
	return x_root == y_root
}

func NewUnionFind(size int) *UnionFind {
	parent := make([]int, size)
	for i := range size {
		parent[i] = i
	}
	uf := UnionFind{size: size, parent: parent, rank: make([]int, size)}
	return &uf
}
