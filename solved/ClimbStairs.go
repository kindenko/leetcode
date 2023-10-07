package main

func ClimbStairs(n int) int {
	fib1 := 0
	fib2 := 1
	total := 0

	for i := 0; i <= n-1; i++ {
		total = fib1 + fib2
		fib1, fib2 = fib2, total

	}

	return total
}
