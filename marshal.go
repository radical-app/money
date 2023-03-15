package money

import (
	"encoding/json"
	"errors"
)

func (m Money) MarshalJSON() ([]byte, error) {
	dto := m.ExtractDTO()

	return json.Marshal(dto)
}

func (m *Money) UnmarshalJSON(data []byte) error {
	dto := DTO{}
	if err := json.Unmarshal(data, &dto); err != nil {
		return err
	}

	if dto.Currency == "" {
		return errors.New("invalid money object: empty currency")
	}

	var err error
	*m, err = dto.ExtractMoney()

	return err
}
