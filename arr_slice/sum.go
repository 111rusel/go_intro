package arrSlice

func Sum(numbers []int) int{
	var sum int
	// for i:= 0; i < 5; i++{
	// 	sum += numbers[i]
	// }
	for _, num := range numbers{
		sum += num
	}
	return sum
}

func SumAll(numbersToSum ...[]int) []int{
	result := []int{}
	for i := range numbersToSum{
		// sum := Sum(numbersToSum[i])
		// result = append(result, sum)
		result = append(result, Sum(numbersToSum[i]))
	}
	return result
}

func SumAllTails(numbersToSum ...[]int) []int{
	result := []int{}
	for i := range numbersToSum{
		if len(numbersToSum[i]) < 2{
			result = append(result, 0)
			continue
		}
		tail := numbersToSum[i][1:]
		result = append(result, Sum(tail))
	}
	return result
}