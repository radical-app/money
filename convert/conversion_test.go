package convert

import (
	"github.com/radicalcompany/money"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestConvertTo(t *testing.T) {
	type args struct {
		obj  money.Money
		rate Rate
	}
	tests := []struct {
		name    string
		args    args
		wantRes money.Money
		wantErr bool
	}{
		{
			name: "eur_to_eur",
			args: args{
				obj:  money.MustForge(100, "EUR"),
				rate: ForgeRate(money.MustGetCurrencyByISOCode("EUR"), money.MustGetCurrencyByISOCode("EUR"), 1),
			},
			wantRes: money.MustForge(100, "EUR"),
			wantErr: false,
		},
		{
			name: "eur_to_usd",
			args: args{
				obj:  money.MustForge(100, "EUR"),
				rate: ForgeRate(money.MustGetCurrencyByISOCode("EUR"), money.MustGetCurrencyByISOCode("USD"), 1.1),
			},
			wantRes: money.MustForge(110, "USD"),
			wantErr: false,
		},
		{
			name: "eur_to_usd_reverse",
			args: args{
				obj:  money.MustForge(100, "EUR"),
				rate: ForgeRate(money.MustGetCurrencyByISOCode("USD"), money.MustGetCurrencyByISOCode("EUR"), 0.91),
			},
			wantRes: money.MustForge(110, "USD"),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotRes, err := ConvertTo(&tt.args.obj, tt.args.rate)
			if (err != nil) != tt.wantErr {
				t.Errorf("ConvertTo() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			assert.NotNil(t, gotRes)
			if !assert.True(t, gotRes.IsEquals(tt.wantRes)) {
				t.Errorf("ConvertTo() gotRes = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}
