package iterationRep

func Repeat(sveta string, count int) string {
	var result string
	for i := 0; i < count; i++ {
		result = result + sveta
	}
	return result
}
