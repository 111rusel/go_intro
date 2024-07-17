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
