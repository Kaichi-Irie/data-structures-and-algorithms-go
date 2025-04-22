package binary

/*
The Peak Finding Problem
Input: A sequence of numbers ⟨a1, a2, . . . , an⟩.
Output: An index i such that ai−1 ≤ ai and ai ≥ ai +1 (define a0 = an+1 = −∞).
If no such index exists, return −1.
*/

func IsPeak(a []int, i int) bool {
	if i < 0 || i >= len(a) {
		return false
	}
	switch l := len(a); {
	case l == 0:
		return false
	case l == 1:
		return true
	case l >= 2 && i == 0:
		return a[0] >= a[1]
	case l >= 2 && i == l-1:
		return a[l-1] >= a[l-2]
	default:
		return a[i-1] <= a[i] && a[i] >= a[i+1]
	}
}

// FindPeak finds a peak in the a[low:high)
func FindPeak(a []int, low int, high int) int {
	if len(a) == 0 || low < 0 || high > len(a) || low >= high {
		return -1
	}

	if high-low == 1 {
		if IsPeak(a, low) {
			return low
		}
		return -1
	}

	mid := (low + high) / 2
	if IsPeak(a, mid) {
		return mid
	}
	if mid > 0 && a[mid] < a[mid-1] {
		return FindPeak(a, low, mid)
	}
	if mid < len(a)-1 && a[mid] < a[mid+1] {
		return FindPeak(a, mid, high)
	}
	return -1
}
