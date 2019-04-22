package fibonacci

// Fibonacci returns list of n fibonacci values
func Fibonacci(n int) []int {
	fibos := make([]int, 0)
	for i := 1; i <= n; i++ {
		switch i {
		case 1:
			fibos = append(fibos, 1)
		case 2:
			fibos = append(fibos, 2)
		default:
			fibos = append(fibos, fibos[i-2]+fibos[i-3])
		}
	}
	return fibos
}

func evenSumFib(max int) int {
	var sum int
	fibos := make([]int, 0)

	for i := 1; true; i++ {
		switch i {
		case 1:
			fibos = append(fibos, 1)
		case 2:
			fibos = append(fibos, 2)
		default:
			fibos = append(fibos, fibos[i-2]+fibos[i-3])
		}
		if fibos[i-1] > max {
			break
		}
		if fibos[i-1]%2 == 0 {
			sum += fibos[i-1]
		}
	}
	return sum
}
