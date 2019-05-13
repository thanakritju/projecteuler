package sumsquare

func sumSquare(n int) int {
	sum1 := 0
	sum2 := 0
	for i := 1; i <= n; i++ {
		sum1 = sum1 + i*i
		sum2 = sum2 + i
	}
	return sum2*sum2 - sum1
}
