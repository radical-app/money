package convert

import (
	"fmt"
	"github.com/radical-app/money"
	"math"
)

type Rate struct {
	Source money.Currency `json:"source"`
	Target money.Currency `json:"target"`
	Rate   float64        `json:"rate"`
}

func ForgeRate(source, target money.Currency, rate float64) Rate {
	return Rate{
		Source: source,
		Target: target,
		Rate:   rate,
	}
}

func ConvertTo(obj *money.Money, rate Rate) (res *money.Money, err error) {
	if rate.Source.IsEquals(obj.Currency) && rate.Target.IsEquals(obj.Currency) {
		return obj, nil
	}

	if rate.Source.IsEquals(obj.Currency) {
		return convertFromSource(obj, rate)
	}
	if rate.Target.IsEquals(obj.Currency) {
		return convertToSource(obj, rate)
	}

	return nil, fmt.Errorf("money currency and rate doesn't match: currency %s, rate source %s, rate target %s", obj.Currency.String(), rate.Source.String(), rate.Target.String())
}


func convertFromSource(obj *money.Money, rate Rate) (res *money.Money, err error) {
	amountFrom := obj.Float()
	toRate := rate.Rate
	centsCount := math.Pow(10, float64(rate.Target.MinorUnit))
	resultAmountInCents := int64(math.Round(amountFrom * toRate * centsCount))
	result, err := money.Forge(resultAmountInCents, rate.Target.Code)
	if err != nil {
		return nil, err
	}
	return &result, err
}


func convertToSource(obj *money.Money, rate Rate) (res *money.Money, err error) {
	amountFrom := obj.Float()
	toRate := rate.Rate
	centsCount := math.Pow(10, float64(rate.Source.MinorUnit))
	resultAmountInCents := int64(math.Round(amountFrom / toRate * centsCount))
	result, err := money.Forge(resultAmountInCents, rate.Source.Code)
	if err != nil {
		return nil, err
	}
	return &result, err
}