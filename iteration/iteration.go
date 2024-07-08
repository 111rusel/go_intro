package iteration

func Repeat(symbol string, count int) string {
	var result string
	for i := 0; i < count; i++ {
		result += symbol
	}
	return result
}
