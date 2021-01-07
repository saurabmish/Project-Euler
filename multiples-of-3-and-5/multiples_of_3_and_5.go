package main

// Multiples return the sum of multiples of 3 and 5 less than 1000
func Multiples() int {
	sum := 0
	for i := 0; i < 1000; i++ {
		if i%3 == 0 || i%5 == 0 {
			sum += i
		}
	}
	return sum
}
