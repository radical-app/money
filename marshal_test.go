package money

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestMoney_MarshalJSON(t *testing.T) {
	tests := []struct {
		name     string
		fields   Money
		wantByte string
		wantErr  bool
	}{
		{"name", MustForge(123, "EUR"), `{"amount":123,"currency":"EUR","symbol":"€","unit":2}`, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &tt.fields

			got, err := json.Marshal(m)
			if (err != nil) != tt.wantErr {
				t.Errorf("Money.MarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(string(got), tt.wantByte) {
				t.Errorf("Money.MarshalJSON() = %v, want %v", string(got), tt.wantByte)
			}

			cmp := &Money{}
			err = json.Unmarshal(got, cmp)
			if (err != nil) != tt.wantErr {
				t.Errorf("Money.Unmarshal() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(cmp, m) {
				t.Errorf("comparing unmarshall= %v, want %v", cmp, m)
			}
		})
	}
}
