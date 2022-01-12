package iteration_2

import "errors"

type Parcel struct {
	width, length, height, price float64
}

var smallParcel = Parcel{
	width:  15,
	length: 10,
	height: 5,
	price:  2.0,
}

var mediumParcel = Parcel{
	width:  20,
	length: 15,
	height: 10,
	price:  3.5,
}

var largeParcel = Parcel{
	width:  25,
	length: 20,
	height: 15,
	price:  7.0,
}

var extraLargeParcel = Parcel{
	width:  30,
	length: 25,
	height: 20,
	price:  10.0,
}

type parcels []Parcel

var allParcelsSizes = parcels{
	smallParcel,
	mediumParcel,
	largeParcel,
	extraLargeParcel,
}

var ParcelTooLargeErr = errors.New("parcel too large for postage")

func (p *Parcel) CalculateCost() error {
	if p.width > extraLargeParcel.width || p.length > extraLargeParcel.length || p.height > extraLargeParcel.height {
		return ParcelTooLargeErr
	}
	for _, parcel := range allParcelsSizes {
		if p.width <= parcel.width && p.length <= parcel.length && p.height <= parcel.height {
			p.price = parcel.price
			break
		}
	}
	return nil
}
