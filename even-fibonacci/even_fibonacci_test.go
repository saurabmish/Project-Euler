package main

import "testing"

func TestEvenFibonacci(t *testing.T) {

	/*
	t.Run("Cutoff greater than sum of first two even numbers", func(t *testing.T) {
		got := EvenFibonacci(4e6)
		want := 4613732

		if got != want {
			t.Errorf("Error in computation; Got: %d Want: %d", got, want)
		}
	})
	*/
	got := EvenFibonacci(4e6)
	want := 4613732

	if got != want {
		t.Errorf("Error in computation; Got: %d Want: %d", got, want)
	}
}
