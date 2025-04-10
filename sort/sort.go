package sort

/*
MergeSort sort the array arr[p, r)
worst time complexity: O(nlogn)
average time complexity: O(nlogn)
best time complexity: O(nlogn)
space complexity:
*/
func MergeSort(arr []int, p int, r int) {
	if r-1 <= p {
		return
	}
	q := (p + r) / 2
	MergeSort(arr, p, q)
	MergeSort(arr, q, r)
	// fmt.Println("Merging", arr[p:q], "and", arr[q:r])
	Merge(arr, p, q, r)
	// fmt.Println("Merged", arr[p:r])
}

/*
Merge merge the two sorted array arr[p, q) and arr[q, r)
worst time complexity: O(n)
average time complexity: O(n)
best time complexity: O(n)
space complexity:
*/
func Merge(arr []int, p int, q int, r int) {
	i := 0
	j := 0
	n1, n2 := q-p, r-q
	X := make([]int, n1)
	Y := make([]int, n2)
	copy(X, arr[p:q])
	copy(Y, arr[q:r])
	for {

		if i >= n1 && j >= n2 {
			break
		}
		if i >= n1 {
			arr[p+i+j] = Y[j]
			j++
			continue
		}
		if j >= n2 {
			arr[p+i+j] = X[i]
			i++
			continue
		}
		if X[i] <= Y[j] {
			arr[p+i+j] = X[i]
			i++
			continue
		}
		if Y[j] <= X[i] {
			arr[p+i+j] = Y[j]
			j++
			continue
		}
	}
}

/*
worst time complexity: O(n^2)
average time complexity: O(n^2)
best time complexity: O(1)

space complexity: O(n)
*/
func InsertionSort(arr []int) {
	n := len(arr)
	if n == 1 {
		return
	}
	for i := 1; i < n; i++ {
		val := arr[i]
		for j := i - 1; j > -1; j-- {
			if arr[j] > val {
				arr[j+1] = arr[j]
				arr[j] = val
			} else {
				break
			}
		}
	}
}
