package tree

type BinaryTree struct {
	Root  int
	Left  *BinaryTree
	Right *BinaryTree
}

func (bt BinaryTree) Depth() int {
	if bt.Left == nil && bt.Right == nil {
		return 1
	}
	var depthLeft, depthRight int
	if bt.Left != nil {
		depthLeft = bt.Left.Depth()
	}
	if bt.Right != nil {
		depthRight = bt.Right.Depth()
	}
	return max(depthLeft, depthRight) + 1
}

func NewBinarySearchTree(sortedArr []int) *BinaryTree {
	if len(sortedArr) == 0 {
		return nil
	}
	mid := len(sortedArr) / 2
	return &BinaryTree{
		Root:  sortedArr[mid],
		Left:  NewBinarySearchTree(sortedArr[:mid]),
		Right: NewBinarySearchTree(sortedArr[mid+1:]),
	}
}

func (bt BinaryTree) find(x int) bool {
	if bt.Root == x {
		return true
	}
	return false
}
