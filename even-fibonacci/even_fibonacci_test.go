package main

import "testing"

func TestEvenFibonacci(t *testing.T) {

	t.Run("Cutoff greater than the sum of first two even numbers", func(t *testing.T) {
		got := EvenFibonacci(4e6)
		want := 4613732

		if got != want {
			t.Errorf("Error in computation; Got: %d Want: %d", got, want)
		}
	})

	// failing test 1
	t.Run("Cutoff less than the sum of first two even numbers", func(t *testing.T) {
		got := EvenFibonacci(4)
		want := 10

		if got != want {
			t.Errorf("Cutoff parameter should be more than 10")
		}
	})

	// failing test 2
	t.Run("Cutoff equal to the sum of first two even numbers", func(t *testing.T) {
		got := EvenFibonacci(10)
		want := 10

		if got != want {
			t.Errorf("Cutoff parameter should be more than 10")
		}
	})
}
