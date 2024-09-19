package search

func lineSearch(s []int, v int) int {
	for i := range s {
		if v == s[i] {
			return i
		}
	}
	return -1
}

func binarySearch(s []int, v int) int {
	low := 0
	high := len(s) - 1
	for low <= high {
		mid := int((low + high) / 2)
		if s[mid] == v {
			return mid
		}
		if s[mid] > v {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return -1
}
