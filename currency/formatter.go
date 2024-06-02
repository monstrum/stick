package currency

import (
	"github.com/bojanz/currency"
)

func NewCurrencyFormatter(locale string) *currency.Formatter {
	return currency.NewFormatter(
		currency.NewLocale(locale),
	)
}
