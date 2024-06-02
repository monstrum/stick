package currency

import (
	"github.com/bojanz/currency"
	"github.com/monstrum/stick"
	"github.com/monstrum/stick/twig/factory"
	"math/big"
)

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
