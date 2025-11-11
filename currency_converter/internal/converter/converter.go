package converter

import (
	"errors"
	"fmt"
)

// Supported returns the list of supported currency codes in UI order.
func Supported() []string {
	return []string{"USD", "GBP", "EUR", "JPY"}
}

// Convert converts amount from currencyFrom to currencyTo using the provided rates map.
// rates should contain USD base rates as provided by Open Exchange Rates.
// For non-USD pairs, we use amount * (rateTo / rateFrom).
func Convert(amount float64, currencyFrom, currencyTo string, rates map[string]float64) (float64, error) {
	if amount < 0 {
		return 0, errors.New("amount must be non-negative")
	}
	if currencyFrom == "" || currencyTo == "" {
		return 0, errors.New("both currencies must be provided")
	}
	if rates == nil {
		return 0, errors.New("rates map is nil")
	}
	
	rateFrom, ok := rates[currencyFrom]
	if !ok {
		return 0, fmt.Errorf("missing rate for %s", currencyFrom)
	}
	rateTo, ok := rates[currencyTo]
	if !ok {
		return 0, fmt.Errorf("missing rate for %s", currencyTo)
	}

	// Direct or cross via USD when base is USD.
	converted := amount * (rateTo / rateFrom)
	return converted, nil
}
