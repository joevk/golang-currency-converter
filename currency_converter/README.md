# Currency Converter (TUI)

A terminal-based currency converter app in Go using Charmbracelet's `huh` for the TUI and Open Exchange Rates for live data.

## Features

- Select from/to currencies (USD, GBP, EUR, JPY)
- Enter an amount to convert
- Live exchange rates via Open Exchange Rates
- Simple, readable TUI form

## Requirements

- Go 1.21+
- Open Exchange Rates API key (free plan works)

## Setup

1. Set your API key environment variable:
   - Windows (PowerShell):
     ```powershell
     $env:OPENEXCHANGERATES_API_KEY = "YOUR_API_KEY"
     ```
2. Build and run:
   ```powershell
   go build ./...
   go run .
   ```

## Notes

- The free plan uses USD as the base currency. Conversion between non-USD pairs uses: `amount * (rateTo / rateFrom)`.
- Handles common errors (network, bad input) with user-friendly messages.

## License

MIT
