package iteration_2

import "testing"

func TestCostCalculator(t *testing.T) {
	cases := []struct {
		Name   string
		Price  float64
		Parcel Parcel
	}{
		{
			Name:  "Small package postage cost should be 2.0",
			Price: 2.0,
			Parcel: Parcel{
				width:  15.0,
				length: 10.0,
				height: 5.0,
			},
		},
		{
			Name:  "Medium package postage cost should be 3.5",
			Price: 3.5,
			Parcel: Parcel{
				width:  20.0,
				length: 15.0,
				height: 10.0,
			},
		},
		{
			Name:  "Large package postage cost should be 7.0",
			Price: 7.0,
			Parcel: Parcel{
				width:  25.0,
				length: 20.0,
				height: 15.0,
			},
		},
		{
			Name:  "Extra Large package postage cost should be 10.0",
			Price: 10.0,
			Parcel: Parcel{
				width:  30,
				length: 25,
				height: 20,
			},
		},
	}

	for _, testCase := range cases {
		t.Run(testCase.Name, func(t *testing.T) {
			testCase.Parcel.CalculateCost()
			got := testCase.Parcel.price
			want := testCase.Price

			assertPrice(t, got, want)
		})
	}

}

func assertPrice(t *testing.T, got float64, want float64) {
	if got != want {
		t.Errorf("got '%e' want '%e'", got, want)
	}
}
