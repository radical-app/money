package money_test

import (
	"fmt"
	"math"
	"testing"

	"github.com/radical-app/money"
	"github.com/stretchr/testify/assert"
)

func TestForge_big(t *testing.T) {
	m, err := money.Forge(math.MaxInt64, "EUR")
	assert.Nil(t, err, err)
	assert.Equal(t, m.Float(), 92233720368547758.07)
}

func TestForge_shouldBrokenOnMonopoliMoney(t *testing.T) {
	_, err := money.Forge(100, "Monopoly")

	assert.NotNil(t, err)
}

func TestForge(t *testing.T) {
	OneEur, err := money.Forge(100, "EUR")
	assert.Nil(t, err)
	assert.Equal(t, OneEur, money.EUR(100))
}

func TestOutputASFloat(t *testing.T) {
	OneEurAndOneCent, err := money.Forge(101, "EUR")
	assert.Nil(t, err)
	assert.Equal(t, OneEurAndOneCent.Float(), 1.01)

	OneEur, err := money.Forge(100, "EUR")
	assert.Nil(t, err)
	assert.Equal(t, OneEur.Float(), float64(1))
}

func TestOutputASFloatZeroDecimals(t *testing.T) {
	amount, err := money.Forge(1011, "VND")
	assert.Nil(t, err)
	assert.Equal(t, float64(1011), amount.Float())

	OneVND, err := money.Forge(1000, "VND")
	assert.Nil(t, err)
	assert.Equal(t, float64(1000), OneVND.Float())
}

func TestOutputASFloatThreeDecimals(t *testing.T) {
	amount, err := money.Forge(1011, "JOD")
	assert.Nil(t, err)
	assert.Equal(t, float64(1.011), amount.Float())

	OneJOD, err := money.Forge(1000, "JOD")
	assert.Nil(t, err)
	assert.Equal(t, float64(1.000), OneJOD.Float())
}

func TestFloat(t *testing.T) {
	OneEurAndOneCent := money.FloatEUR(1.01)

	assert.Equal(t, OneEurAndOneCent.Int64(), int64(101))
	assert.Equal(t, OneEurAndOneCent.Float(), 1.01)
}

func TestAmountAsString(t *testing.T) {
	OneEurAndOneCent := money.FloatEUR(1.01)

	assert.Equal(t, OneEurAndOneCent.Int64(), int64(101))
	assert.Equal(t, OneEurAndOneCent.AmountAsString(), "1.01")
}

func TestOutputAsStringZeroDecimals(t *testing.T) {
	amount, err := money.Forge(1011, "VND")
	assert.Nil(t, err)
	assert.Equal(t, "1011", amount.AmountAsString())

	OneVND, err := money.Forge(1000, "VND")
	assert.Nil(t, err)
	assert.Equal(t, "1000", OneVND.AmountAsString())
}

func TestOutputAStringThreeDecimals(t *testing.T) {
	amount, err := money.Forge(1011, "JOD")
	assert.Nil(t, err)
	assert.Equal(t, "1.011", amount.AmountAsString())

	OneJOD, err := money.Forge(1000, "JOD")
	assert.Nil(t, err)
	assert.Equal(t, "1.000", OneJOD.AmountAsString())
}

func TestMoney_IsEquals(t *testing.T) {
	OneEurAndOneCent := money.FloatEUR(1.01)

	assert.True(t, OneEurAndOneCent.IsEquals(money.EUR(101)))

	assert.False(t, OneEurAndOneCent.IsEquals(money.GBP(101)))
}

func TestMoney_PercentageOff(t *testing.T) {
	OneEurAndOneCent := money.FloatEUR(100.09)

	pOff := OneEurAndOneCent.PercentOff(20)

	assert.Equal(t, pOff.Int64(), int64(2002))
	assert.Equal(t, pOff.Float(), 20.02)

	assert.True(t, OneEurAndOneCent.IsEquals(money.EUR(10009)))
}

func TestMoney_PercentageOffFloat(t *testing.T) {
	amount := money.FloatEUR(100.09)

	pOff := amount.PercentOffFloat(20.5)

	assert.Equal(t, pOff.Int64(), int64(2052))
	assert.Equal(t, pOff.Float(), 20.52)

	assert.True(t, amount.IsEquals(money.EUR(10009)))
}

func TestMoney_CentsValue(t *testing.T) {
	type fields struct {
		Amount      int64
		AmountFloat float64
		Currency    string
	}
	tests := []struct {
		name       string
		fields     fields
		wantValueI int64
		wantI      int
		wantErr    bool
	}{
		{"test", fields{0, 1234.56789, "EUR"}, 1234, 57, false},

		{"test", fields{0, 1234.5, "EUR"}, 1234, 5, false},

		{"test", fields{0, 123.45, "EUR"}, 123, 45, false},

		{"test", fields{0, 12.345, "EUR"}, 12, 35, false},

		{"test", fields{0, 1.2345, "EUR"}, 1, 23, false},

		{"test", fields{0, 12.345, "JPY"}, 12, 0, false},

		{"test float", fields{12345, 0, "EUR"}, 123, 45, false},

		{"test float", fields{12345, 0, "JPY"}, 12345, 0, false},

		{"test float", fields{12345, 0, "CLF"}, 0, 12345, false},

		{"test float", fields{123, 0, "CLF"}, 0, 123, false},

		{"test float", fields{12, 0, "EUR"}, 0, 12, false},

		{"test error", fields{0, 12.12312312, "EUR"}, 12, 12, false},

		{"test error", fields{0, 0, "EUR"}, 0, 0, false},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("name: %s %v", tt.name, tt.fields), func(t *testing.T) {
			var m money.Money
			var err error
			if tt.fields.AmountFloat == 0 {
				m, err = money.Forge(tt.fields.Amount, tt.fields.Currency)
			} else {
				m, err = money.ForgeFloat(tt.fields.AmountFloat, tt.fields.Currency)
			}

			if (err != nil) != tt.wantErr {
				t.Errorf("Money.CentsValue() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			valueI, gotI, err := m.SplitAmountAndCents()
			if (err != nil) != tt.wantErr {
				t.Errorf("Money.CentsValue() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotI != tt.wantI {
				t.Errorf("Money.CentsValue() = %v, want %v", gotI, tt.wantI)
			}

			if valueI != tt.wantValueI {
				t.Errorf("Money.CentsValue() = %v, wantValueI %v", valueI, tt.wantValueI)
			}
		})
	}
}
