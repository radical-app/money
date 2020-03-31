package money

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

func ParseWithFallback(s string, fallbackCurr Currency) (m Money, err error) {
	m = EUR(0)
	code := DefaultCurrencyCode
	curr := fallbackCurr
	if !fallbackCurr.IsValid() {
		curr, err = CurrencyByISOCode(code)
		if err != nil {
			return m, err
		}
	}
	m = ForgeWithCurrency(0, curr)

	if s == "" {
		return m, errors.New("empty string")
	}

	ss := strings.Fields(s)
	if len(ss) != 1 && len(ss) != 2 {
		return m, fmt.Errorf("money field should be like `EUR 123` given %v", s)
	}

	amountAsInt, err := strconv.ParseInt(ss[len(ss)-1], 10, 64)
	if err != nil {
		return m, err
	}

	// the first group is the curr code iso code
	// use default Currency
	if len(ss) == 2 {
		code = ss[0]
		curr, err = CurrencyByISOCode(code)
		if err != nil {
			return m, err
		}
	}

	return ForgeWithCurrency(amountAsInt, curr), err
}

// Parse Create a money object by a string like "EUR 123" "CurrencyCode Int64" or "Int64"
func Parse(s string) (m Money, err error) {
	return ParseWithFallback(s, Currency{})
}
