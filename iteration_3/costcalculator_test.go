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
			Name:  "Parcel 15x10x5 should be considered 'Small' and price should be 2.0",
			Price: 2.0,
			Parcel: Parcel{
				width:  15.0,
				length: 10.0,
				height: 5.0,
			},
			expectedErr: nil,
		},
		{
			Name:  "Parcel 20x15x10 should be considered 'Medium' and price should be 3.5",
			Price: 3.5,
			Parcel: Parcel{
				width:  20.0,
				length: 15.0,
				height: 10.0,
			},
			expectedErr: nil,
		},
		{
			Name:  "Parcel 25x20x15 should be considered 'Large' and price should be 7.0",
			Price: 7.0,
			Parcel: Parcel{
				width:  25.0,
				length: 20.0,
				height: 15.0,
			},
			expectedErr: nil,
		},
		{
			Name:  "Parcel 15x1x1 should be considered 'Large' and price should be 7.0",
			Price: 7.0,
			Parcel: Parcel{
				width:  25.0,
				length: 1.0,
				height: 1.0,
			},
			expectedErr: nil,
		},
		{
			Name:  "Parcel 1x10x1 should be considered 'Large' and price should be 7.0",
			Price: 7.0,
			Parcel: Parcel{
				width:  1.0,
				length: 20.0,
				height: 1.0,
			},
			expectedErr: nil,
		},
		{
			Name:  "Parcel 1x1x5 should be considered 'Large' and price should be 7.0",
			Price: 7.0,
			Parcel: Parcel{
				width:  1.0,
				length: 1.0,
				height: 15.0,
			},
			expectedErr: nil,
		},
		{
			Name:  "Parcel 30x25x20 should be considered 'Extra Large' and price should be 10.0",
			Price: 10.0,
			Parcel: Parcel{
				width:  30,
				length: 25,
				height: 20,
			},
			expectedErr: nil,
		},
		{
			Name: "Parcel 31x25x20 should be too large 'ParcelTooLargeErr' error should be returned",
			Parcel: Parcel{
				width:  31,
				length: 25,
				height: 20,
			},
			expectedErr: ParcelTooLargeErr,
		},
		{
			Name: "Parcel 30x26x20 should be too large 'ParcelTooLargeErr' error should be returned",
			Parcel: Parcel{
				width:  30,
				length: 26,
				height: 20,
			},
			expectedErr: ParcelTooLargeErr,
		},
		{
			Name: "Parcel 30x25x21 should be too large 'ParcelTooLargeErr' error should be returned",
			Parcel: Parcel{
				width:  30,
				length: 25,
				height: 21,
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
		t.Errorf("got '%s' want '%s'", got, want)
	}

}

func assertPrice(t *testing.T, got float64, want float64) {
	t.Helper()
	if got != want {
		t.Errorf("got '%e' want '%e'", got, want)
	}
}
