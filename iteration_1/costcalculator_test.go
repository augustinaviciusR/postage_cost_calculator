package iteration_1

import "testing"

func TestCostCalculator(t *testing.T) {
	got := CalculateCost(10, 20, 4)
	want := 2.0

	if got != want {
		t.Errorf("got '%e' want '%e'", got, want)
	}
}
