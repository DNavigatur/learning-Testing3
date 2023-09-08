package main

import "testing"

func TestAverage(t *testing.T) {
	xi := []float64{2, 4, 7, 14}
	got := Average(xi)
	if got != 6.75 {
		t.Error("Expected", 6.5, "Got", got)
	}
}
