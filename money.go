package money

import (
	"errors"
	"fmt"
	"math"
	"strconv"

	"database/sql/driver"
)

type Amount int64

func (a Amount) Int64() int64 {
	return int64(a)
}

type Money struct {
	Amount   Amount   `json:"amount"`
	Currency Currency `json:"currency"`
}

func (m Money) Int64() int64 {
	return int64(m.Amount)
}

// Forge
// amount   int64   A positive integer in cents (or 0 for a free plan)
//                  representing how much to charge on a recurring basis.
// currCode string Three-letter ISO currency code, in lowercase
func Forge(amount int64, currCode string) (m Money, err error) {
	c, err := CurrencyByISOCode(currCode)
	if err != nil {
		return
	}
	m = ForgeWithCurrency(amount, c)
	return
}

// Forge
// amount   float64 A positive float64 in cents (or 0 for a free plan)
//          representing how much to charge on a recurring basis.
// currCode string  Three-letter ISO currency code, in lowercase
func ForgeFloat(amount float64, currCode string) (m Money, err error) {
	c, err := CurrencyByISOCode(currCode)
	if err != nil {
		return m, err
	}

	m = ForgeFloatWithCurrency(amount, c)
	return m, err
}

func MustForgeFloat(amountFloat float64, currCode string) Money {
	m, err := Forge(0, currCode)
	if err != nil {
		panic(err)
	}

	return ForgeFloatWithCurrency(amountFloat, m.Currency)
}

func ForgeFloatWithCurrency(amountFloat float64, c Currency) Money {
	fd := float64(c.GetCents())
	amount := int64(math.Round(fd * amountFloat))

	return ForgeWithCurrency(amount, c)
}

// MustForge Forge or panic
// amount   int    A positive integer in cents (or 0 for a free plan)
//                 representing how much to charge on a recurring basis.
// currCode string Three-letter ISO currency code, in lowercase
func MustForge(amount int64, currCode string) Money {
	m, err := Forge(amount, currCode)
	if err != nil {
		panic(err)
	}

	return m
}

// ForgeWithCurrency
// amount   int      A positive integer in cents (or 0 for a free plan) representing how much to charge on a recurring basis.
// currency Currency The currency Value Object
func ForgeWithCurrency(amount int64, c Currency) Money {
	var err error
	if !c.IsValid() {
		c, err = CurrencyByISOCode(DefaultCurrencyCode)
		if err != nil {
			panic(err)
		}
	}

	return Money{Amount: Amount(amount), Currency: c}
}

func (m Money) IsZero() bool {
	return m.Amount == 0
}

func (m Money) Float() float64 {
	if m.IsZero() {
		return 0
	}
	d := m.DigitsAsCents()
	return float64(m.Amount) / float64(d)
}

func (m Money) DigitsAsCents() int {
	return m.Currency.GetCents()
}

func (m Money) IsEquals(cmp Money) bool {
	return m.Amount == cmp.Amount && m.Currency.IsEquals(cmp.Currency)
}

func (m Money) PercentOff(perc int) Money {
	div := float64(perc) / 100
	return ForgeFloatWithCurrency(m.Float()*div, m.Currency)
}

func (m Money) Add(addendum Money) (s Money, err error) {
	if !m.Currency.IsEquals(addendum.Currency) {
		return s, errors.New(fmt.Sprint("Can't compare or use math with different currency", m.Currency, s.Currency))
	}

	return ForgeWithCurrency(addendum.Amount.Int64()+m.Amount.Int64(), m.Currency), err
}

func (m Money) Subtract(subtrahend Money) (s Money, err error) {
	if !m.Currency.IsEquals(subtrahend.Currency) {
		return s, errors.New(fmt.Sprint("Can't compare or use math with different currency", m.Currency, s.Currency))
	}

	return ForgeWithCurrency(m.Amount.Int64()-subtrahend.Amount.Int64(), m.Currency), err
}

func (m Money) SplitAmountAndCents() (i int64, cents int, err error) {
	f := m.Float()
	digitsCount := m.Currency.MinorUnit

	i = int64(f)
	if f-float64(i) <= 0 {
		return i, cents, err
	}

	cutAt := 2 + digitsCount
	ct := fmt.Sprintf("%d", cutAt)
	decPartStr := fmt.Sprintf("%."+ct+"g", f-float64(i))

	if len(decPartStr) <= 2 {
		return i, cents, errors.New("can't convert to decimal " + decPartStr)
	}

	if len(decPartStr) > cutAt {
		cents, err = strconv.Atoi(decPartStr[2:cutAt])
		return i, cents, err
	}

	cents, err = strconv.Atoi(decPartStr[2:])
	return i, cents, err
}

func Scan(value interface{}, curr string) (m Money, err error) {
	if v, ok := value.(int64); ok {
		mm, err := Forge(v, curr)
		return mm, err
	}
	return m, errors.New(fmt.Sprintf("impossible to get int64 the value from %v", value))
}

func (m *Money) ScanInt64(value interface{}) error {
	if value == nil {
		return nil
	}
	code := DefaultCurrencyCode
	if m.Currency.IsValid() {
		code = m.Currency.Code
	}
	mm, err := Scan(value, code)
	if err != nil {
		return err
	}
	*m = mm
	return nil
}

func (m *Money) Scan(value interface{}) error {
	if value == nil {
		return nil
	}

	if v, ok := value.(string); ok {
		mm, err := ParseWithFallback(v, m.Currency)
		if err != nil {
			return err
		}
		*m = mm
		return err
	}

	if v, ok := value.(int64); ok {
		mm := ForgeWithCurrency(v, m.Currency)
		err := mm.ScanInt64(value)
		if err != nil {
			return err
		}
		*m = mm
		return err
	}

	if v, ok := value.(float64); ok {
		code := DefaultCurrencyCode
		if m.Currency.IsValid() {
			code = m.Currency.Code
		}
		mm, err := ForgeFloat(v, code)
		if err != nil {
			return err
		}
		*m = mm
		return nil
	}

	return errors.New(fmt.Sprintf("can't convert given %v", value))
}

// Value implements the driver Valuer interface.
func (m *Money) Value() (driver.Value, error) {
	return m.Int64(), nil
}

// Value implements the driver Valuer interface.
func (m *Money) String() string {
	return fmt.Sprintf("%s %d", m.Currency.String(), m.Int64())
}
