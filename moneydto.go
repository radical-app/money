package money

type DTO struct {
	Amount   int64  `json:"amount"`
	Currency string `json:"currency"`
	Symbol   string `json:"symbol"`
	Cents    int    `json:"cents"`
}

func (d DTO) ExtractMoney() (m Money, err error) {
	return Forge(d.Amount, d.Currency)
}

func (m Money) ExtractDTO() DTO {
	return DTO{m.Amount.Int64(),
		m.Currency.Code,
		m.Currency.Symbol,
		m.Currency.GetCents(),
	}
}
