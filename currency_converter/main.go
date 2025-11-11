package main

import (
	"fmt"
	"log"
	"os"

	"currency_converter/internal/api"
	"currency_converter/internal/converter"
	"currency_converter/internal/tui"
)

// main orchestrates the application loop.
// Steps:
// 1. Fetch rates once at start (could be cached for a session).
// 2. Show TUI form for user input.
// 3. Perform conversion and display result.
// 4. Ask if they want another conversion (simple prompt).\n//    For simplicity we reuse the form for new conversions; rates stay same.
func main() {
	client := api.NewClient()
	data, err := client.FetchRates()
	if err != nil {
		log.Fatalf("Failed to fetch rates: %v", err)
	}

	fmt.Println("Live rates loaded (base USD). Supported: USD, GBP, EUR, JPY")

	for {
		res, err := tui.RunForm()
		if err != nil {
			// Handle user abort or form error.
			fmt.Fprintf(os.Stderr, "Form error: %v\n", err)
			return
		}

		amountConv, err := converter.Convert(res.Amount, res.From, res.To, data.Rates)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Conversion error: %v\n", err)
			continue
		}

		fmt.Printf("\nResult: %.4f %s -> %.4f %s\n", res.Amount, res.From, amountConv, res.To)

		// Ask user if they want another conversion.
		var again string
		fmt.Print("Another conversion? (y/N): ")
		fmt.Scanln(&again)
		if again != "y" && again != "Y" {
			break
		}
	}

	fmt.Println("Goodbye!")
}
