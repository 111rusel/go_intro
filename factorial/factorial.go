package factorial

func factorial(n int) int {
	// result := 1
	// for i := 1; i <= n; i++ {
	// 	result = result * i
	// }
	// return result
	if n == 0 {
		return 1
	}
	return factorial(n-1) * n
}

/*
5! = 1 * 2 * 3 * 4 * 5 = 4! * 5                 5! - это факториал пяти
4! = 1 * 2 * 3 * 4 = 3! * 4
3! = 1 * 2 * 3 = 2! * 3
2! = 1 * 2 = 1! * 2
1! = 1 * 1 = 0! * 1
0! = 1 - закон математики
*/
