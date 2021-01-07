package main

// EvenFibonacci returns the sum of even Fibonacci terms up to a limit
func EvenFibonacci(cutoff int) int {
	firstEven := 2
	secondEven := 8
	evenSum := firstEven + secondEven
	for evenSum < cutoff {
		evenFib := 4 * secondEven + firstEven
		evenSum += evenFib
		firstEven, secondEven = secondEven, evenFib
	}
	return evenSum
}
