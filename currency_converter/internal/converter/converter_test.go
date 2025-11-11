package converter

import "testing"

func TestConvert(t *testing.T) {
	rates := map[string]float64{
		"USD": 1.0,
		"EUR": 0.9,
		"GBP": 0.8,
		"JPY": 150.0,
	}

	// USD -> EUR
	if got, err := Convert(100, "USD", "EUR", rates); err != nil || got < 89.99 || got > 90.01 {
		t.Fatalf("USD->EUR: got=%v err=%v", got, err)
	}
	// EUR -> USD
	if got, err := Convert(90, "EUR", "USD", rates); err != nil || got < 99.99 || got > 100.01 {
		t.Fatalf("EUR->USD: got=%v err=%v", got, err)
	}
	// EUR -> GBP cross via USD
	if got, err := Convert(100, "EUR", "GBP", rates); err != nil {
		t.Fatalf("EUR->GBP err=%v", err)
	} else {
		expected := 100 * (rates["GBP"] / rates["EUR"]) // 100 * (0.8/0.9) = 88.888...
		if got < expected-0.0001 || got > expected+0.0001 {
			t.Fatalf("EUR->GBP: got=%v expected=%v", got, expected)
		}
	}
}
