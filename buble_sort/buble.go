package bubleSort

func BubleSort(s []int) []int {
	var isSliceChanged bool = true
	for isSliceChanged {
		isSliceChanged = false
		for i := 0; i < len(s)-1; i++ {
			if s[i] > s[i+1] {
				s[i], s[i+1] = s[i+1], s[i]
				isSliceChanged = true
			}
		}
	}
	return s
}
