package util

//supported currencies

const (
	USD = "USD"
	EUR = "EUR"
	CAD = "CAD"
	NGN = "NGN"
)

// IsSupportedCurrency returns true if supported and false otherwise
func IsSupportedCurrency(currency string) bool {
	switch currency {
	case USD, EUR, CAD, NGN:
		return true
	}
	return false
}
