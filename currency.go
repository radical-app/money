package money

import (
	"errors"
	"math"
)

var DefaultCurrencyCode = "EUR"

// GetCurrencyByCode gets the currency object by currency ISO code
func CurrencyByISOCode(code string) (currency Currency, err error) {
	if c, ok := currencies[code]; ok {
		return c, nil
	}

	return currency, errors.New("Currency not found")
}

// GetCurrencyByCode gets the currency object by currency ISO code
func MustGetCurrencyByISOCode(code string) (currency Currency) {
	currency, err := CurrencyByISOCode(code)
	if err != nil {
		panic(err)
	}

	return currency
}

// Currency codes
// https://www.iso.org/iso-4217-currency-codes.html
type Currency struct {
	Code                 string `json:"currency"`
	MinorUnit            int    `json:"unit"`
	Symbol               string `json:"symbol"`
	ShowCodeNextToSymbol bool
}

func (c Currency) IsZeroDigitsAfterDecimalSeparator() bool {
	return c.MinorUnit == 0
}

func (c Currency) IsValid() bool {
	c, err := CurrencyByISOCode(c.Code)
	if err != nil {
		return false
	}

	return c.Code != ""
}

func (c Currency) GetCents() int {
	if c.MinorUnit < 0 {
		return 1
	}
	ce := math.Pow(10, float64(c.MinorUnit))
	return int(ce)
}

func (c Currency) String() string {
	return c.Code
}

func (c Currency) IsEquals(cmp Currency) bool {
	return c.Code == cmp.Code && c.MinorUnit == cmp.MinorUnit
}

var currencies = map[string]Currency{
	"AED": {Code: "AED", MinorUnit: 2, Symbol: ".\u062f.\u0625", ShowCodeNextToSymbol: true},
	"AFN": {Code: "AFN", MinorUnit: 2, Symbol: "\u060b", ShowCodeNextToSymbol: false},
	"ALL": {Code: "ALL", MinorUnit: 2, Symbol: "Lek", ShowCodeNextToSymbol: false},
	"AMD": {Code: "AMD", MinorUnit: 2, Symbol: "\u0564\u0580.", ShowCodeNextToSymbol: false},
	"ANG": {Code: "ANG", MinorUnit: 2, Symbol: "\u0192", ShowCodeNextToSymbol: true},
	"AOA": {Code: "AOA", MinorUnit: 2, Symbol: "Kz", ShowCodeNextToSymbol: false},
	"ARS": {Code: "ARS", MinorUnit: 2, Symbol: "$", ShowCodeNextToSymbol: true},
	"AUD": {Code: "AUD", MinorUnit: 2, Symbol: "A$", ShowCodeNextToSymbol: false},
	"AWG": {Code: "AWG", MinorUnit: 2, Symbol: "\u0192", ShowCodeNextToSymbol: true},
	"AZN": {Code: "AZN", MinorUnit: 2, Symbol: "\u20bc", ShowCodeNextToSymbol: false},
	"BAM": {Code: "BAM", MinorUnit: 2, Symbol: "KM", ShowCodeNextToSymbol: false},
	"BBD": {Code: "BBD", MinorUnit: 2, Symbol: "Bds$", ShowCodeNextToSymbol: false},
	"BDT": {Code: "BDT", MinorUnit: 2, Symbol: "\u09f3", ShowCodeNextToSymbol: false},
	"BGN": {Code: "BGN", MinorUnit: 2, Symbol: "\u043b\u0432", ShowCodeNextToSymbol: false},
	"BHD": {Code: "BHD", MinorUnit: 3, Symbol: ".\u062f.\u0628", ShowCodeNextToSymbol: false},
	"BIF": {Code: "BIF", MinorUnit: 0, Symbol: "FBu", ShowCodeNextToSymbol: false},
	"BMD": {Code: "BMD", MinorUnit: 2, Symbol: "BD$", ShowCodeNextToSymbol: false},
	"BND": {Code: "BND", MinorUnit: 2, Symbol: "", ShowCodeNextToSymbol: false},
	"BOB": {Code: "BOB", MinorUnit: 2, Symbol: "Bs.", ShowCodeNextToSymbol: false},
	"BRL": {Code: "BRL", MinorUnit: 2, Symbol: "R$", ShowCodeNextToSymbol: false},
	"BSD": {Code: "BSD", MinorUnit: 2, Symbol: "", ShowCodeNextToSymbol: false},
	"BTN": {Code: "BTN", MinorUnit: 2, Symbol: "Nu.", ShowCodeNextToSymbol: false},
	"BWP": {Code: "BWP", MinorUnit: 2, Symbol: "P", ShowCodeNextToSymbol: true},
	"BYN": {Code: "BYN", MinorUnit: 2, Symbol: "Br", ShowCodeNextToSymbol: false},
	"BZD": {Code: "BZD", MinorUnit: 2, Symbol: "BZ$", ShowCodeNextToSymbol: false},
	"CAD": {Code: "CAD", MinorUnit: 2, Symbol: "CAD$", ShowCodeNextToSymbol: false},
	"CDF": {Code: "CDF", MinorUnit: 2, Symbol: "FC", ShowCodeNextToSymbol: false},
	"CHF": {Code: "CHF", MinorUnit: 2, Symbol: "", ShowCodeNextToSymbol: false},
	"CLF": {Code: "CLF", MinorUnit: 5, Symbol: "UF", ShowCodeNextToSymbol: false},
	"CLP": {Code: "CLP", MinorUnit: 0, Symbol: "CLP$", ShowCodeNextToSymbol: false},
	"CNY": {Code: "CNY", MinorUnit: 2, Symbol: "\u5143", ShowCodeNextToSymbol: false},
	"COP": {Code: "COP", MinorUnit: 2, Symbol: "COP$", ShowCodeNextToSymbol: false},
	"CRC": {Code: "CRC", MinorUnit: 2, Symbol: "\u20a1", ShowCodeNextToSymbol: true},
	"CUC": {Code: "CUC", MinorUnit: 2, Symbol: "CUC$", ShowCodeNextToSymbol: false},
	"CUP": {Code: "CUP", MinorUnit: 2, Symbol: "$MN", ShowCodeNextToSymbol: false},
	"CVE": {Code: "CVE", MinorUnit: 2, Symbol: "Esc", ShowCodeNextToSymbol: false},
	"CZK": {Code: "CZK", MinorUnit: 2, Symbol: "K\u010d", ShowCodeNextToSymbol: false},
	"DJF": {Code: "DJF", MinorUnit: 0, Symbol: "Fdj", ShowCodeNextToSymbol: false},
	"DKK": {Code: "DKK", MinorUnit: 2, Symbol: "kr", ShowCodeNextToSymbol: true},
	"DOP": {Code: "DOP", MinorUnit: 2, Symbol: "RD$", ShowCodeNextToSymbol: false},
	"DZD": {Code: "DZD", MinorUnit: 2, Symbol: ".\u062f.\u062c", ShowCodeNextToSymbol: false},
	"EGP": {Code: "EGP", MinorUnit: 2, Symbol: "\u00a3", ShowCodeNextToSymbol: true},
	"ERN": {Code: "ERN", MinorUnit: 2, Symbol: "Nfk", ShowCodeNextToSymbol: false},
	"ETB": {Code: "ETB", MinorUnit: 2, Symbol: "Br", ShowCodeNextToSymbol: false},
	"EUR": {Code: "EUR", MinorUnit: 2, Symbol: "\u20ac", ShowCodeNextToSymbol: false},
	"FJD": {Code: "FJD", MinorUnit: 2, Symbol: "FJ$", ShowCodeNextToSymbol: false},
	"FKP": {Code: "FKP", MinorUnit: 2, Symbol: "\u00a3", ShowCodeNextToSymbol: true},
	"GBP": {Code: "GBP", MinorUnit: 2, Symbol: "\u00a3", ShowCodeNextToSymbol: false},
	"GEL": {Code: "GEL", MinorUnit: 2, Symbol: "\u10da", ShowCodeNextToSymbol: false},
	"GGP": {Code: "GGP", MinorUnit: 2, Symbol: "\u00a3", ShowCodeNextToSymbol: false},
	"GHS": {Code: "GHS", MinorUnit: 2, Symbol: "\u20b5", ShowCodeNextToSymbol: false},
	"GIP": {Code: "GIP", MinorUnit: 2, Symbol: "\u00a3", ShowCodeNextToSymbol: true},
	"GMD": {Code: "GMD", MinorUnit: 2, Symbol: "D", ShowCodeNextToSymbol: false},
	"GNF": {Code: "GNF", MinorUnit: 0, Symbol: "FG", ShowCodeNextToSymbol: false},
	"GTQ": {Code: "GTQ", MinorUnit: 2, Symbol: "Q", ShowCodeNextToSymbol: false},
	"GYD": {Code: "GYD", MinorUnit: 2, Symbol: "G$", ShowCodeNextToSymbol: false},
	"HKD": {Code: "HKD", MinorUnit: 2, Symbol: "HK$", ShowCodeNextToSymbol: false},
	"HNL": {Code: "HNL", MinorUnit: 2, Symbol: "L", ShowCodeNextToSymbol: true},
	"HRK": {Code: "HRK", MinorUnit: 2, Symbol: "kn", ShowCodeNextToSymbol: false},
	"HTG": {Code: "HTG", MinorUnit: 2, Symbol: "G", ShowCodeNextToSymbol: false},
	"HUF": {Code: "HUF", MinorUnit: 0, Symbol: "Ft", ShowCodeNextToSymbol: false},
	"IDR": {Code: "IDR", MinorUnit: 2, Symbol: "Rp", ShowCodeNextToSymbol: false},
	"ILS": {Code: "ILS", MinorUnit: 2, Symbol: "\u20aa", ShowCodeNextToSymbol: false},
	"IMP": {Code: "IMP", MinorUnit: 2, Symbol: "\u00a3", ShowCodeNextToSymbol: true},
	"INR": {Code: "INR", MinorUnit: 2, Symbol: "\u20b9", ShowCodeNextToSymbol: false},
	"IQD": {Code: "IQD", MinorUnit: 3, Symbol: ".\u062f.\u0639", ShowCodeNextToSymbol: false},
	"IRR": {Code: "IRR", MinorUnit: 2, Symbol: "\ufdfc", ShowCodeNextToSymbol: true},
	"ISK": {Code: "ISK", MinorUnit: 0, Symbol: "kr", ShowCodeNextToSymbol: true},
	"JEP": {Code: "JEP", MinorUnit: 2, Symbol: "\u00a3", ShowCodeNextToSymbol: true},
	"JMD": {Code: "JMD", MinorUnit: 2, Symbol: "J$", ShowCodeNextToSymbol: false},
	"JOD": {Code: "JOD", MinorUnit: 3, Symbol: ".\u062f.\u0625", ShowCodeNextToSymbol: true},
	"JPY": {Code: "JPY", MinorUnit: 0, Symbol: "\u00a5", ShowCodeNextToSymbol: false},
	"KES": {Code: "KES", MinorUnit: 2, Symbol: "KSh", ShowCodeNextToSymbol: false},
	"KGS": {Code: "KGS", MinorUnit: 2, Symbol: "\u0441\u043e\u043c", ShowCodeNextToSymbol: false},
	"KHR": {Code: "KHR", MinorUnit: 2, Symbol: "\u17db", ShowCodeNextToSymbol: false},
	"KMF": {Code: "KMF", MinorUnit: 0, Symbol: "CF", ShowCodeNextToSymbol: false},
	"KPW": {Code: "KPW", MinorUnit: 0, Symbol: "\u20a9", ShowCodeNextToSymbol: true},
	"KRW": {Code: "KRW", MinorUnit: 0, Symbol: "\u20a9", ShowCodeNextToSymbol: true},
	"KWD": {Code: "KWD", MinorUnit: 3, Symbol: ".\u062f.\u0643", ShowCodeNextToSymbol: false},
	"KYD": {Code: "KYD", MinorUnit: 2, Symbol: "CI$", ShowCodeNextToSymbol: false},
	"KZT": {Code: "KZT", MinorUnit: 2, Symbol: "\u20b8", ShowCodeNextToSymbol: false},
	"LAK": {Code: "LAK", MinorUnit: 2, Symbol: "\u20ad", ShowCodeNextToSymbol: false},
	"LBP": {Code: "LBP", MinorUnit: 2, Symbol: "\u00a3", ShowCodeNextToSymbol: true},
	"LKR": {Code: "LKR", MinorUnit: 2, Symbol: "\u20a8", ShowCodeNextToSymbol: true},
	"LRD": {Code: "LRD", MinorUnit: 2, Symbol: "L$", ShowCodeNextToSymbol: false},
	"LSL": {Code: "LSL", MinorUnit: 2, Symbol: "L", ShowCodeNextToSymbol: true},
	"LYD": {Code: "LYD", MinorUnit: 3, Symbol: ".\u062f.\u0644", ShowCodeNextToSymbol: false},
	"MAD": {Code: "MAD", MinorUnit: 2, Symbol: ".\u062f.\u0645", ShowCodeNextToSymbol: false},
	"MDL": {Code: "MDL", MinorUnit: 2, Symbol: "lei", ShowCodeNextToSymbol: true},
	"MGA": {Code: "MGA", MinorUnit: 0, Symbol: "Ar", ShowCodeNextToSymbol: false},
	"MKD": {Code: "MKD", MinorUnit: 2, Symbol: "\u0434\u0435\u043d", ShowCodeNextToSymbol: false},
	"MMK": {Code: "MMK", MinorUnit: 2, Symbol: "K", ShowCodeNextToSymbol: true},
	"MNT": {Code: "MNT", MinorUnit: 2, Symbol: "\u20ae", ShowCodeNextToSymbol: false},
	"MOP": {Code: "MOP", MinorUnit: 2, Symbol: "P", ShowCodeNextToSymbol: true},
	"MRO": {Code: "MRO", MinorUnit: 0, Symbol: "UM", ShowCodeNextToSymbol: false},
	"MUR": {Code: "MUR", MinorUnit: 2, Symbol: "\u20a8", ShowCodeNextToSymbol: true},
	"MVR": {Code: "MVR", MinorUnit: 2, Symbol: "MVR", ShowCodeNextToSymbol: false},
	"MWK": {Code: "MWK", MinorUnit: 2, Symbol: "MK", ShowCodeNextToSymbol: false},
	"MXN": {Code: "MXN", MinorUnit: 2, Symbol: "Mex$", ShowCodeNextToSymbol: false},
	"MYR": {Code: "MYR", MinorUnit: 2, Symbol: "RM", ShowCodeNextToSymbol: false},
	"MZN": {Code: "MZN", MinorUnit: 2, Symbol: "MT", ShowCodeNextToSymbol: false},
	"NAD": {Code: "NAD", MinorUnit: 2, Symbol: "N$", ShowCodeNextToSymbol: false},
	"NGN": {Code: "NGN", MinorUnit: 2, Symbol: "\u20a6", ShowCodeNextToSymbol: false},
	"NIO": {Code: "NIO", MinorUnit: 2, Symbol: "C$", ShowCodeNextToSymbol: false},
	"NOK": {Code: "NOK", MinorUnit: 2, Symbol: "kr", ShowCodeNextToSymbol: true},
	"NPR": {Code: "NPR", MinorUnit: 2, Symbol: "\u20a8", ShowCodeNextToSymbol: true},
	"NZD": {Code: "NZD", MinorUnit: 2, Symbol: "NZ$", ShowCodeNextToSymbol: false},
	"OMR": {Code: "OMR", MinorUnit: 3, Symbol: "\ufdfc", ShowCodeNextToSymbol: true},
	"PAB": {Code: "PAB", MinorUnit: 2, Symbol: "B/.", ShowCodeNextToSymbol: false},
	"PEN": {Code: "PEN", MinorUnit: 2, Symbol: "S/", ShowCodeNextToSymbol: false},
	"PGK": {Code: "PGK", MinorUnit: 2, Symbol: "K", ShowCodeNextToSymbol: true},
	"PHP": {Code: "PHP", MinorUnit: 2, Symbol: "\u20b1", ShowCodeNextToSymbol: false},
	"PKR": {Code: "PKR", MinorUnit: 2, Symbol: "\u20a8", ShowCodeNextToSymbol: true},
	"PLN": {Code: "PLN", MinorUnit: 2, Symbol: "z\u0142", ShowCodeNextToSymbol: false},
	"PYG": {Code: "PYG", MinorUnit: 0, Symbol: "Gs", ShowCodeNextToSymbol: false},
	"QAR": {Code: "QAR", MinorUnit: 2, Symbol: "\ufdfc", ShowCodeNextToSymbol: true},
	"RON": {Code: "RON", MinorUnit: 2, Symbol: "lei", ShowCodeNextToSymbol: true},
	"RSD": {Code: "RSD", MinorUnit: 2, Symbol: "\u0414\u0438\u043d.", ShowCodeNextToSymbol: false},
	"RUB": {Code: "RUB", MinorUnit: 2, Symbol: "\u20bd", ShowCodeNextToSymbol: false},
	"RWF": {Code: "RWF", MinorUnit: 0, Symbol: "FRw", ShowCodeNextToSymbol: false},
	"SAR": {Code: "SAR", MinorUnit: 2, Symbol: "\ufdfc", ShowCodeNextToSymbol: true},
	"SBD": {Code: "SBD", MinorUnit: 2, Symbol: "SI$", ShowCodeNextToSymbol: false},
	"SCR": {Code: "SCR", MinorUnit: 2, Symbol: "\u20a8", ShowCodeNextToSymbol: true},
	"SDG": {Code: "SDG", MinorUnit: 2, Symbol: "\u00a3", ShowCodeNextToSymbol: true},
	"SEK": {Code: "SEK", MinorUnit: 2, Symbol: "kr", ShowCodeNextToSymbol: true},
	"SGD": {Code: "SGD", MinorUnit: 2, Symbol: "S$", ShowCodeNextToSymbol: false},
	"SHP": {Code: "SHP", MinorUnit: 2, Symbol: "\u00a3", ShowCodeNextToSymbol: true},
	"SLL": {Code: "SLL", MinorUnit: 2, Symbol: "Le", ShowCodeNextToSymbol: false},
	"SOS": {Code: "SOS", MinorUnit: 2, Symbol: "Sh", ShowCodeNextToSymbol: false},
	"SRD": {Code: "SRD", MinorUnit: 2, Symbol: "", ShowCodeNextToSymbol: false},
	"SSP": {Code: "SSP", MinorUnit: 2, Symbol: "\u00a3", ShowCodeNextToSymbol: true},
	"STD": {Code: "STD", MinorUnit: 2, Symbol: "Db", ShowCodeNextToSymbol: false},
	"SVC": {Code: "SVC", MinorUnit: 2, Symbol: "\u20a1", ShowCodeNextToSymbol: true},
	"SYP": {Code: "SYP", MinorUnit: 2, Symbol: "\u00a3", ShowCodeNextToSymbol: true},
	"SZL": {Code: "SZL", MinorUnit: 2, Symbol: "\u00a3", ShowCodeNextToSymbol: true},
	"THB": {Code: "THB", MinorUnit: 2, Symbol: "\u0e3f", ShowCodeNextToSymbol: false},
	"TJS": {Code: "TJS", MinorUnit: 2, Symbol: "SM", ShowCodeNextToSymbol: false},
	"TMT": {Code: "TMT", MinorUnit: 2, Symbol: "T", ShowCodeNextToSymbol: true},
	"TND": {Code: "TND", MinorUnit: 3, Symbol: ".\u062f.\u062a", ShowCodeNextToSymbol: false},
	"TOP": {Code: "TOP", MinorUnit: 2, Symbol: "T$", ShowCodeNextToSymbol: false},
	"TRY": {Code: "TRY", MinorUnit: 2, Symbol: "\u20ba", ShowCodeNextToSymbol: false},
	"TTD": {Code: "TTD", MinorUnit: 2, Symbol: "TT$", ShowCodeNextToSymbol: false},
	"TWD": {Code: "TWD", MinorUnit: 0, Symbol: "NT$", ShowCodeNextToSymbol: false},
	"TZS": {Code: "TZS", MinorUnit: 0, Symbol: "TSh", ShowCodeNextToSymbol: false},
	"UAH": {Code: "UAH", MinorUnit: 2, Symbol: "\u20b4", ShowCodeNextToSymbol: false},
	"UGX": {Code: "UGX", MinorUnit: 0, Symbol: "USh", ShowCodeNextToSymbol: false},
	"USD": {Code: "USD", MinorUnit: 2, Symbol: "$", ShowCodeNextToSymbol: false},
	"UYU": {Code: "UYU", MinorUnit: 2, Symbol: "$U", ShowCodeNextToSymbol: false},
	"UZS": {Code: "UZS", MinorUnit: 2, Symbol: "so\u2019m", ShowCodeNextToSymbol: false},
	"VES": {Code: "VES", MinorUnit: 2, Symbol: "Bs.S", ShowCodeNextToSymbol: false},
	"VEF": {Code: "VEF", MinorUnit: 2, Symbol: "Bs.F", ShowCodeNextToSymbol: false},
	"VND": {Code: "VND", MinorUnit: 0, Symbol: "\u20ab", ShowCodeNextToSymbol: false},
	"VUV": {Code: "VUV", MinorUnit: 0, Symbol: "Vt", ShowCodeNextToSymbol: false},
	"WST": {Code: "WST", MinorUnit: 2, Symbol: "T", ShowCodeNextToSymbol: true},
	"XAF": {Code: "XAF", MinorUnit: 0, Symbol: "Fr", ShowCodeNextToSymbol: true},
	"XOF": {Code: "XOF", MinorUnit: 0, Symbol: "Fr", ShowCodeNextToSymbol: true},
	"XPF": {Code: "XPF", MinorUnit: 0, Symbol: "Fr", ShowCodeNextToSymbol: true},
	"XCD": {Code: "XCD", MinorUnit: 2, Symbol: "EC$", ShowCodeNextToSymbol: false},
	"YER": {Code: "YER", MinorUnit: 2, Symbol: "\ufdfc", ShowCodeNextToSymbol: true},
	"ZAR": {Code: "ZAR", MinorUnit: 2, Symbol: "R", ShowCodeNextToSymbol: false},
	"ZMW": {Code: "ZMW", MinorUnit: 2, Symbol: "ZK", ShowCodeNextToSymbol: false},
	"ZWD": {Code: "ZWD", MinorUnit: 2, Symbol: "Z$", ShowCodeNextToSymbol: false},
}
