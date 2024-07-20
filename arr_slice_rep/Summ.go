package arrSliceRep

func Summ(numbers [5]int) int {
	var result int
	for i := range numbers {
		result = result + numbers[i]
	}
	return result
}

func SummSlice(numbers []int) int {
	var result int
	for i := range numbers {
		result = result + numbers[i]
	}
	return result
}

func SummAll(slices ...[]int) []int {
	result := []int{}
	for i := range slices {
		summ := SummSlice(slices[i])
		result = append(result, summ)
	}
	return result
}

func SumAllTails(slices ...[]int) []int {
	result := []int{}
	for _, s := range slices {
		if len(s) < 2 {
			result = append(result, 0)
		} else {
			result = append(result, SummSlice(s[1:]))
		}
	}
	return result
}
