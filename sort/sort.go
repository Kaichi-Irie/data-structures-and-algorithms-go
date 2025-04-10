package sort

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
