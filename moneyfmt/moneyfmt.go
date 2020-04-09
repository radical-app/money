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
	_, cents, err := m.SplitAmountAndCents()
	if err != nil {
		return formatted, err
	}
	floatAmountAsString := p.Sprintf("%f", m.Float())
	maxI := lastIndexOfCommaOrDot(floatAmountAsString)
	if cents == 0 && maxI <= len(floatAmountAsString) {
		return floatAmountAsString[0:maxI], nil
	}
	centsAsString := fmt.Sprintf("%d", cents)

	return floatAmountAsString[0 : maxI+1+len(centsAsString)], nil
}

func MustDisplayAmount(m money.Money, locale string) (formatted string) {
	formatted, err := DisplayAmount(m, locale)
	if err != nil {
		panic(err)
	}

	return formatted
}

// Display Symbol
func Display(m money.Money, locale string) (formatted string, err error) {
	a, err := DisplayAmount(m, locale)
	if err != nil {
		return a, err
	}
	return fmt.Sprintf("%s %s", m.Currency.Symbol, a), err
}

func MustDisplay(m money.Money, locale string) (formatted string) {
	formatted, err := Display(m, locale)
	if err != nil {
		panic(err)
	}

	return formatted
}

// Display Symbol
func DisplayISO(m money.Money, locale string) (formatted string, err error) {
	a, err := DisplayAmount(m, locale)
	if err != nil {
		return a, err
	}
	return fmt.Sprintf("%s %s", m.Currency.Code, a), err
}

func MustDisplayISO(m money.Money, locale string) (formatted string) {
	formatted, err := DisplayISO(m, locale)
	if err != nil {
		panic(err)
	}

	return formatted
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
