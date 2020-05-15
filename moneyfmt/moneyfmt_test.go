package moneyfmt_test

import (
	"github.com/stretchr/testify/assert"
	"testing"

	"golang.org/x/text/currency"
	"golang.org/x/text/language"
	"golang.org/x/text/message"

	"github.com/radical-app/money"
	"github.com/radical-app/money/moneyfmt"
)

func TestDisplay(t *testing.T) {
	tests := []struct {
		name          string
		args          money.Money
		wantFormatted string
		wantErr       bool
	}{
		{"ru", money.MustForge(123400, "EUR"), "€ 1 234", false},
		{"ru", money.MustForge(123456, "EUR"), "€ 1 234,56", false},

		{"it", money.MustForge(123456, "EUR"), "€ 1.234,56", false},
		{"it", money.MustForge(123400, "EUR"), "€ 1.234", false},
		{"en", money.MustForge(123456, "EUR"), "€ 1,234.56", false},

		{"it", money.MustForge(123400, "EUR"), "€ 1.234", false},

		{"jp", money.MustForge(123456, "EUR"), "€ 1,234.56", false},
		{"zh", money.MustForge(123456, "EUR"), "€ 1,234.56", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotFormatted, err := moneyfmt.Display(tt.args, tt.name)
			if (err != nil) != tt.wantErr {
				t.Errorf("Display() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotFormatted != tt.wantFormatted {
				t.Errorf("Display() = %v, want %v", gotFormatted, tt.wantFormatted)
			}
		})
	}
}

func TestMustDisplay(t *testing.T) {
	tests := []struct {
		name          string
		args          money.Money
		wantFormatted string
	}{
		{"ru", money.MustForge(123400, "EUR"), "€ 1 234"},
		{"ru", money.MustForge(123456, "EUR"), "€ 1 234,56"},

		{"it", money.MustForge(123456, "EUR"), "€ 1.234,56"},
		{"it", money.MustForge(123400, "EUR"), "€ 1.234"},
		{"en", money.MustForge(123456, "EUR"), "€ 1,234.56"},

		{"it", money.MustForge(123400, "EUR"), "€ 1.234"},

		{"jp", money.MustForge(123456, "EUR"), "€ 1,234.56"},
		{"zh", money.MustForge(123456, "EUR"), "€ 1,234.56"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.NotPanics(t, func() {
				gotFormatted := moneyfmt.MustDisplay(tt.args, tt.name)
				if gotFormatted != tt.wantFormatted {
					t.Errorf("Display() = %v, want %v", gotFormatted, tt.wantFormatted)
				}
			})
		})
	}
}

func TestByLocale(t *testing.T) {
	v := currency.NarrowSymbol
	tag := language.Make("it")

	output := v(currency.USD.Amount(12348.20))

	pIt := message.NewPrinter(tag)
	assertE := "$ 12.348,20"
	out := pIt.Sprint(output)
	if out != assertE {
		t.Skip("Waiting  Currency on golang works as aspected ", assertE, out)
		return
	}
}
