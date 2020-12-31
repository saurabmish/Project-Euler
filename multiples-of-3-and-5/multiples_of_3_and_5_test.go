package main

import "testing"

func TestMultiples(t *testing.T) {
	got := Multiples()
	want := 233168

	if got != want {
		t.Errorf("Got: %q Want: %q", got, want)
	}
}
