package moneyfmt

import (
	"fmt"
	"strings"

	"github.com/radicalcompany/money"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

func DisplayAmount(m money.Money, locale string) (formatted string, err error) {
	tag := language.Make(locale)
	p := message.NewPrinter(tag)
	_, c, err := m.SplitAmountAndCents()
	if err != nil {
		return formatted, err
	}
	o := p.Sprintf("%f", m.Float())
	maxI := lastIndexOfCommaOrDot(o)
	if c == 0 && maxI <= len(o) {
		return o[0:maxI], nil
	}
	centsAsString := fmt.Sprintf("%d", c)

	return o[0 : maxI+1+len(centsAsString)], nil
}

// Display Symbol
func Display(m money.Money, locale string) (formatted string, err error) {
	a, err := DisplayAmount(m, locale)
	if err != nil {
		return a, err
	}
	return fmt.Sprintf("%s %s", m.Currency.Symbol, a), err
}

// Display Symbol
func DisplayISO(m money.Money, locale string) (formatted string, err error) {
	a, err := DisplayAmount(m, locale)
	if err != nil {
		return a, err
	}
	return fmt.Sprintf("%s %s", m.Currency.Code, a), err
}

func lastIndexOfCommaOrDot(o string) int {
	iDot := strings.LastIndex(o, ".")
	iComma := strings.LastIndex(o, ",")
	maxI := iComma
	if iDot > iComma {
		maxI = iDot
	}
	return maxI
}
