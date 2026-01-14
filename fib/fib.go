package fib

// N returns the nth Fibonacci number (0-indexed)
func N(n int) int {
	if n <= 0 {
		return 0
	}
	if n == 1 {
		return 1
	}

	a, b := 0, 1
	for i := 2; i <= n; i++ {
		a, b = b, a+b
	}
	return b
}

// Sequence returns the first n Fibonacci numbers
func Sequence(n int) []int {
	if n <= 0 {
		return []int{}
	}

	result := make([]int, n)
	for i := 0; i < n; i++ {
		result[i] = N(i)
	}
	return result
}
