package currency

import (
	"github.com/bojanz/currency"
	"github.com/monstrum/stick"
	"github.com/monstrum/stick/twig/factory"
	"math/big"
)

// AddFormatCurrencyFilter adds a format_currency filter to the filters map.
// this filter doesn't have full support for all the options of the PHP version.
// The locale is used to determine the currency symbol and format.
// The currency code is used to determine the currency symbol.
// The currency code is optional and defaults to EUR.
// The value can be a string, int, int32, int64, or big.Int.
func AddFormatCurrencyFilter(locale string) factory.AppendFilterFn {
	return func(filters map[string]stick.Filter) {
		filters["format_currency"] = filterFormatCurrencyFn(locale)
	}
}

func filterFormatCurrencyFn(locale string) stick.Filter {
	formatter := NewCurrencyFormatter(locale)
	return func(ctx stick.Context, val stick.Value, args ...stick.Value) stick.Value {
		currencyCode := "EUR"
		if len(args) > 0 {
			currencyCode = stick.CoerceString(args[0])
		}

		var amount currency.Amount
		var err error
		switch v := val.(type) {
		case string:
			amount, err = currency.NewAmount(v, currencyCode)
		case int64:
		case int32:
		case int:
			amount, err = currency.NewAmountFromInt64(int64(v), currencyCode)
		case big.Int:
			amount, err = currency.NewAmountFromBigInt(&v, currencyCode)
		}

		if err != nil {
			return val
		}

		return stick.Value(formatter.Format(amount))
	}
}
