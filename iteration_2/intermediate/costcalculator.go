package iteration_2

const smallPackagePostagePrice = 2.0
const mediumPackagePostagePrice = 3.5
const largePackagePostagePrice = 7.0
const extraLargePackagePostagePrice = 10

func CalculateCost(width, length, height float64) float64 {
	if width <= 15 && length <= 10 && height <= 5 {
		return smallPackagePostagePrice
	}
	if width <= 20 && length <= 15 && height <= 10 {
		return mediumPackagePostagePrice
	}
	if width <= 25 && length <= 20 && height <= 15 {
		return largePackagePostagePrice
	}
	return extraLargePackagePostagePrice
}
