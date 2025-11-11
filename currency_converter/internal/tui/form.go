package tui

import (
	"fmt"
	"strconv"

	"currency_converter/internal/converter"

	"github.com/charmbracelet/huh"
)

// FormResult holds user selections from the TUI.
type FormResult struct {
	From   string
	To     string
	Amount float64
}

// RunForm displays the interactive form and returns the user input.
// It loops on validation errors until valid input is provided or the user aborts.
func RunForm() (*FormResult, error) {
	currs := converter.Supported()
	var from, to, amountStr string

	form := huh.NewForm(
		huh.NewGroup(
			huh.NewSelect[string]().Title("Convert FROM").Options(stringOptions(currs)...).Value(&from),
			huh.NewSelect[string]().Title("Convert TO").Options(stringOptions(currs)...).Value(&to),
			huh.NewInput().Title("Amount").Value(&amountStr).Placeholder("e.g. 100.50"),
		),
	).WithTheme(huh.ThemeBase())

	if err := form.Run(); err != nil {
		return nil, err
	}

	amt, err := strconv.ParseFloat(amountStr, 64)
	if err != nil {
		return nil, fmt.Errorf("invalid amount: %w", err)
	}

	return &FormResult{From: from, To: to, Amount: amt}, nil
}

func stringOptions(values []string) []huh.Option[string] {
	opts := make([]huh.Option[string], 0, len(values))
	for _, v := range values {
		opts = append(opts, huh.NewOption[string](v, v))
	}
	return opts
}
