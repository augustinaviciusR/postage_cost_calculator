package iteration_2

const smallPackagePostagePrice = 2.0
const mediumPackagePostagePrice = 3.5
const largePackagePostagePrice = 7.0
const extraLargePackagePostagePrice = 10

type Parcel struct {
	width, length, height, price float64
}
type parcels []Parcel

var allParcelsSizes = parcels{
	{
		width:  15,
		length: 10,
		height: 5,
		price:  smallPackagePostagePrice,
	},
	{
		width:  20,
		length: 15,
		height: 10,
		price:  mediumPackagePostagePrice,
	},
	{
		width:  25,
		length: 20,
		height: 15,
		price:  largePackagePostagePrice,
	},
}

func (p *Parcel) CalculateCost() {
	p.price = extraLargePackagePostagePrice
	for _, parcel := range allParcelsSizes {
		if p.width <= parcel.width && p.length <= parcel.length && p.height <= parcel.height {
			p.price = parcel.price
			break
		}
	}
}
