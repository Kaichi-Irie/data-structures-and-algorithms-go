package binary

/*
Searches for the leftmost index of k in a sorted array A.
If k is not found, returns -1.
*/
func BisectLeft(A []int, k int) int {
	if len(A) == 0 {
		return -1
	}
	L := 0
	R := len(A) - 1
	if k <= A[L] {
		return 0
	}
	if k > A[R] {
		return len(A)
	}
	for R-L > 1 {
		mid := (L + R) / 2
		if A[mid] < k {
			L = mid
		} else {
			R = mid
		}
	}
	return R
}

/*
Searches for the leftmost index of k in a sorted array A.
If k is not found, returns -1.
*/
func BisectRight(A []int, k int) int {
	if len(A) == 0 {
		return -1
	}
	L := 0
	R := len(A) - 1
	if k < A[L] {
		return 0
	}
	if k >= A[R] {
		return len(A)
	}
	for R-L > 1 {
		mid := (L + R) / 2
		if A[mid] <= k {
			L = mid
		} else {
			R = mid
		}
	}
	return R
}
