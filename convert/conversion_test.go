package convert

import (
	"github.com/radical-app/money"
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
		{
			name: "gbp_to_usd",
			args: args{
				obj:  money.MustForge(100, "GBP"),
				rate: ForgeRate(money.MustGetCurrencyByISOCode("USD"), money.MustGetCurrencyByISOCode("EUR"), 2),
			},
			wantErr: true,
		},
		{
			name: "vnd_to_eur",
			args: args{
				obj:  money.MustForge(100000, "VND"),
				rate: ForgeRate(money.MustGetCurrencyByISOCode("EUR"), money.MustGetCurrencyByISOCode("VND"), 25258.410459),
			},
			wantRes: money.MustForge(396, "EUR"),
			wantErr: false,
		},
		{
			name: "eur_to_vnd",
			args: args{
				obj:  money.MustForge(100, "EUR"),
				rate: ForgeRate(money.MustGetCurrencyByISOCode("EUR"), money.MustGetCurrencyByISOCode("VND"), 25258.410459),
			},
			wantRes: money.MustForge(25258, "VND"),
			wantErr: false,
		},
		{
			name: "eur_to_tnd",
			args: args{
				obj:  money.MustForge(100, "EUR"),
				rate: ForgeRate(money.MustGetCurrencyByISOCode("EUR"), money.MustGetCurrencyByISOCode("TND"), 3.148775),
			},
			wantRes: money.MustForge(3149, "TND"),
			wantErr: false,
		},
		{
			name: "tnd_to_eur",
			args: args{
				obj:  money.MustForge(3149, "TND"),
				rate: ForgeRate(money.MustGetCurrencyByISOCode("EUR"), money.MustGetCurrencyByISOCode("TND"), 3.148775),
			},
			wantRes: money.MustForge(100, "EUR"),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotRes, err := ConvertTo(&tt.args.obj, tt.args.rate)
			if tt.wantErr {
				assert.NotNil(t, err)
			}
			if gotRes != nil {
				if !assert.True(t, gotRes.IsEquals(tt.wantRes)) {
					t.Errorf("ConvertTo() result = %v, want %v", gotRes, tt.wantRes)
				}
			}
		})
	}
}
