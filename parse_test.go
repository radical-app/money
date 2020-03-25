package money

import (
	"reflect"
	"testing"
)

func TestForgeByString(t *testing.T) {
	tests := []struct {
		input   string
		wantM   Money
		wantErr bool
	}{
		{"123 123", EUR(0), true},
		{"a", EUR(0), true},
		{"0", EUR(0), false},
		{"USD 12.12", EUR(0), true},
		{"USD 12,12", EUR(0), true},
		{"  EUR   123 today", EUR(0), true},

		{"", EUR(0), true},
		{"EUR 123", EUR(123), false},
		{"  EUR   123  ", EUR(123), false},
		{"USD 21312312312312323", USD(21312312312312323), false},
		{"3324", EUR(3324), false},
	}
	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			gotM, err := Parse(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("ForgeByString() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotM, tt.wantM) {
				t.Errorf("ForgeByString() = %v, want %v", gotM, tt.wantM)
			}
		})
	}
}
