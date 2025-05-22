package binary

/*
The Searching Problem
Input: A sequence of numbers A= ⟨a1, a2, . . . , an⟩ and a value v.
Output: An index i such that v = A[i ], or NIL if v does not appear in A.
*/

/*
BinarySearch performs a binary search on an array A[n1:n2] for the value k.
BinarySearch(A, 0, len(A), k) solves the searching problem.
It returns true if k is found in the array, and false otherwise.
*/
func BinarySearch(A []int, n1 int, n2 int, k int) bool {
	if len(A) == 0 || n1 < 0 || n2 > len(A) || n1 >= n2 {
		return false
	}
	if n2-n1 == 1 {
		return A[n1] == k
	}

	mid := (n1 + n2) / 2
	isInLeft := BinarySearch(A, n1, mid, k)
	if isInLeft {
		return true
	}
	isInRight := BinarySearch(A, mid, n2, k)
	return isInRight
}
