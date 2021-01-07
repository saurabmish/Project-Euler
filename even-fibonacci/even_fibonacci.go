package main

// cutoff value must be greater than the sum of first two even numbers
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
