package iteration_2

import "testing"

func TestCostCalculator(t *testing.T) {
	t.Run("Small package postage cost should be 2.0", func(t *testing.T) {
		got := CalculateCost(15, 10, 5)
		want := 2.0

		if got != want {
			t.Errorf("got '%e' want '%e'", got, want)
		}

	})
	t.Run("Medium package postage cost should be 3.5", func(t *testing.T) {
		got := CalculateCost(20, 15, 10)
		want := 3.5

		if got != want {
			t.Errorf("got '%e' want '%e'", got, want)
		}

	})
	t.Run("Large package postage cost should be 7.0", func(t *testing.T) {
		got := CalculateCost(25, 20, 15)
		want := 7.0

		if got != want {
			t.Errorf("got '%e' want '%e'", got, want)
		}

	})
	t.Run("Extra Large package postage cost should be 10.0", func(t *testing.T) {
		got := CalculateCost(30, 25, 20)
		want := 10.0

		if got != want {
			t.Errorf("got '%e' want '%e'", got, want)
		}

	})
}
