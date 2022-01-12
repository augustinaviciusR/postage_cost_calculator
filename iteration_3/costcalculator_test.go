package iteration_2

import "testing"

func TestCostCalculator(t *testing.T) {
	cases := []struct {
		Name        string
		Price       float64
		Parcel      Parcel
		expectedErr error
	}{
		{
			Name:  "Small package postage cost should be 2.0",
			Price: 2.0,
			Parcel: Parcel{
				width:  15.0,
				length: 10.0,
				height: 5.0,
			},
			expectedErr: nil,
		},
		{
			Name:  "Medium package postage cost should be 3.5",
			Price: 3.5,
			Parcel: Parcel{
				width:  20.0,
				length: 15.0,
				height: 10.0,
			},
			expectedErr: nil,
		},
		{
			Name:  "Large package postage cost should be 7.0",
			Price: 7.0,
			Parcel: Parcel{
				width:  25.0,
				length: 20.0,
				height: 15.0,
			},
			expectedErr: nil,
		},
		{
			Name:  "Extra Large package postage cost should be 10.0",
			Price: 10.0,
			Parcel: Parcel{
				width:  30,
				length: 25,
				height: 20,
			},
			expectedErr: nil,
		},
		{
			Name:  "Too Large package postage error should be returned",
			Price: 0.0,
			Parcel: Parcel{
				width:  31,
				length: 26,
				height: 25,
			},
			expectedErr: ParcelTooLargeErr,
		},
	}

	for _, testCase := range cases {
		t.Run(testCase.Name, func(t *testing.T) {
			err := testCase.Parcel.CalculateCost()
			got := testCase.Parcel.price
			want := testCase.Price

			assertPrice(t, got, want)
			assertError(t, err, testCase.expectedErr)
		})
	}

}

func assertError(t *testing.T, got, want error) {
	t.Helper()
	if got != want {
		t.Errorf("got '%e' want '%e'", got, want)
	}

}

func assertPrice(t *testing.T, got float64, want float64) {
	t.Helper()
	if got != want {
		t.Errorf("got '%e' want '%e'", got, want)
	}
}
