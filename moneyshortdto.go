package money

type ShortDTO string

func (d ShortDTO) ExtractMoney(s string) (m Money, err error) {
	return ParseWithFallback(s, Currency{})
}

func (m Money) ExtractShortDTO() ShortDTO {
	return ShortDTO(m.String())
}
