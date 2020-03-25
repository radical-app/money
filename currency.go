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

// Currency codes
// https://www.iso.org/iso-4217-currency-codes.html
type Currency struct {
	Code                 string `json:"currency"`
	Digits               int    `json:"cents"`
	Symbol               string `json:"symbol"`
	ShowCodeNextToSymbol bool
}

func (c Currency) IsZeroDigitsAfterDecimalSeparator() bool {
	return c.Digits == 0
}

func (c Currency) IsValid() bool {
	c, err := CurrencyByISOCode(c.Code)
	if err != nil {
		return false
	}

	return c.Code != ""
}

func (c Currency) GetCents() int {
	if c.Digits < 0 {
		return 1
	}
	ce := math.Pow(10, float64(c.Digits))
	return int(ce)
}

func (c Currency) String() string {
	return c.Code
}

func (c Currency) IsEquals(cmp Currency) bool {
	return c.Code == cmp.Code && c.Digits == cmp.Digits
}

var currencies = map[string]Currency{
	"AED": {Code: "AED", Digits: 2, Symbol: ".\u062f.\u0625", ShowCodeNextToSymbol: true},
	"AFN": {Code: "AFN", Digits: 2, Symbol: "\u060b", ShowCodeNextToSymbol: false},
	"ALL": {Code: "ALL", Digits: 2, Symbol: "Lek", ShowCodeNextToSymbol: false},
	"AMD": {Code: "AMD", Digits: 2, Symbol: "\u0564\u0580.", ShowCodeNextToSymbol: false},
	"ANG": {Code: "ANG", Digits: 2, Symbol: "\u0192", ShowCodeNextToSymbol: true},
	"AOA": {Code: "AOA", Digits: 2, Symbol: "Kz", ShowCodeNextToSymbol: false},
	"ARS": {Code: "ARS", Digits: 2, Symbol: "$", ShowCodeNextToSymbol: true},
	"AUD": {Code: "AUD", Digits: 2, Symbol: "A$", ShowCodeNextToSymbol: false},
	"AWG": {Code: "AWG", Digits: 2, Symbol: "\u0192", ShowCodeNextToSymbol: true},
	"AZN": {Code: "AZN", Digits: 2, Symbol: "\u20bc", ShowCodeNextToSymbol: false},
	"BAM": {Code: "BAM", Digits: 2, Symbol: "KM", ShowCodeNextToSymbol: false},
	"BBD": {Code: "BBD", Digits: 2, Symbol: "Bds$", ShowCodeNextToSymbol: false},
	"BDT": {Code: "BDT", Digits: 2, Symbol: "\u09f3", ShowCodeNextToSymbol: false},
	"BGN": {Code: "BGN", Digits: 2, Symbol: "\u043b\u0432", ShowCodeNextToSymbol: false},
	"BHD": {Code: "BHD", Digits: 3, Symbol: ".\u062f.\u0628", ShowCodeNextToSymbol: false},
	"BIF": {Code: "BIF", Digits: 0, Symbol: "FBu", ShowCodeNextToSymbol: false},
	"BMD": {Code: "BMD", Digits: 2, Symbol: "BD$", ShowCodeNextToSymbol: false},
	"BND": {Code: "BND", Digits: 2, Symbol: "", ShowCodeNextToSymbol: false},
	"BOB": {Code: "BOB", Digits: 2, Symbol: "Bs.", ShowCodeNextToSymbol: false},
	"BRL": {Code: "BRL", Digits: 2, Symbol: "R$", ShowCodeNextToSymbol: false},
	"BSD": {Code: "BSD", Digits: 2, Symbol: "", ShowCodeNextToSymbol: false},
	"BTN": {Code: "BTN", Digits: 2, Symbol: "Nu.", ShowCodeNextToSymbol: false},
	"BWP": {Code: "BWP", Digits: 2, Symbol: "P", ShowCodeNextToSymbol: true},
	"BYN": {Code: "BYN", Digits: 2, Symbol: "Br", ShowCodeNextToSymbol: false},
	"BZD": {Code: "BZD", Digits: 2, Symbol: "BZ$", ShowCodeNextToSymbol: false},
	"CAD": {Code: "CAD", Digits: 2, Symbol: "CAD$", ShowCodeNextToSymbol: false},
	"CDF": {Code: "CDF", Digits: 2, Symbol: "FC", ShowCodeNextToSymbol: false},
	"CHF": {Code: "CHF", Digits: 2, Symbol: "", ShowCodeNextToSymbol: false},
	"CLF": {Code: "CLF", Digits: 5, Symbol: "UF", ShowCodeNextToSymbol: false},
	"CLP": {Code: "CLP", Digits: 0, Symbol: "CLP$", ShowCodeNextToSymbol: false},
	"CNY": {Code: "CNY", Digits: 2, Symbol: "\u5143", ShowCodeNextToSymbol: false},
	"COP": {Code: "COP", Digits: 2, Symbol: "COP$", ShowCodeNextToSymbol: false},
	"CRC": {Code: "CRC", Digits: 2, Symbol: "\u20a1", ShowCodeNextToSymbol: true},
	"CUC": {Code: "CUC", Digits: 2, Symbol: "CUC$", ShowCodeNextToSymbol: false},
	"CUP": {Code: "CUP", Digits: 2, Symbol: "$MN", ShowCodeNextToSymbol: false},
	"CVE": {Code: "CVE", Digits: 2, Symbol: "Esc", ShowCodeNextToSymbol: false},
	"CZK": {Code: "CZK", Digits: 2, Symbol: "K\u010d", ShowCodeNextToSymbol: false},
	"DJF": {Code: "DJF", Digits: 0, Symbol: "Fdj", ShowCodeNextToSymbol: false},
	"DKK": {Code: "DKK", Digits: 2, Symbol: "kr", ShowCodeNextToSymbol: true},
	"DOP": {Code: "DOP", Digits: 2, Symbol: "RD$", ShowCodeNextToSymbol: false},
	"DZD": {Code: "DZD", Digits: 2, Symbol: ".\u062f.\u062c", ShowCodeNextToSymbol: false},
	"EGP": {Code: "EGP", Digits: 2, Symbol: "\u00a3", ShowCodeNextToSymbol: true},
	"ERN": {Code: "ERN", Digits: 2, Symbol: "Nfk", ShowCodeNextToSymbol: false},
	"ETB": {Code: "ETB", Digits: 2, Symbol: "Br", ShowCodeNextToSymbol: false},
	"EUR": {Code: "EUR", Digits: 2, Symbol: "\u20ac", ShowCodeNextToSymbol: false},
	"FJD": {Code: "FJD", Digits: 2, Symbol: "FJ$", ShowCodeNextToSymbol: false},
	"FKP": {Code: "FKP", Digits: 2, Symbol: "\u00a3", ShowCodeNextToSymbol: true},
	"GBP": {Code: "GBP", Digits: 2, Symbol: "\u00a3", ShowCodeNextToSymbol: false},
	"GEL": {Code: "GEL", Digits: 2, Symbol: "\u10da", ShowCodeNextToSymbol: false},
	"GGP": {Code: "GGP", Digits: 2, Symbol: "\u00a3", ShowCodeNextToSymbol: false},
	"GHS": {Code: "GHS", Digits: 2, Symbol: "\u20b5", ShowCodeNextToSymbol: false},
	"GIP": {Code: "GIP", Digits: 2, Symbol: "\u00a3", ShowCodeNextToSymbol: true},
	"GMD": {Code: "GMD", Digits: 2, Symbol: "D", ShowCodeNextToSymbol: false},
	"GNF": {Code: "GNF", Digits: 0, Symbol: "FG", ShowCodeNextToSymbol: false},
	"GTQ": {Code: "GTQ", Digits: 2, Symbol: "Q", ShowCodeNextToSymbol: false},
	"GYD": {Code: "GYD", Digits: 2, Symbol: "G$", ShowCodeNextToSymbol: false},
	"HKD": {Code: "HKD", Digits: 2, Symbol: "HK$", ShowCodeNextToSymbol: false},
	"HNL": {Code: "HNL", Digits: 2, Symbol: "L", ShowCodeNextToSymbol: true},
	"HRK": {Code: "HRK", Digits: 2, Symbol: "kn", ShowCodeNextToSymbol: false},
	"HTG": {Code: "HTG", Digits: 2, Symbol: "G", ShowCodeNextToSymbol: false},
	"HUF": {Code: "HUF", Digits: 0, Symbol: "Ft", ShowCodeNextToSymbol: false},
	"IDR": {Code: "IDR", Digits: 2, Symbol: "Rp", ShowCodeNextToSymbol: false},
	"ILS": {Code: "ILS", Digits: 2, Symbol: "\u20aa", ShowCodeNextToSymbol: false},
	"IMP": {Code: "IMP", Digits: 2, Symbol: "\u00a3", ShowCodeNextToSymbol: true},
	"INR": {Code: "INR", Digits: 2, Symbol: "\u20b9", ShowCodeNextToSymbol: false},
	"IQD": {Code: "IQD", Digits: 3, Symbol: ".\u062f.\u0639", ShowCodeNextToSymbol: false},
	"IRR": {Code: "IRR", Digits: 2, Symbol: "\ufdfc", ShowCodeNextToSymbol: true},
	"ISK": {Code: "ISK", Digits: 0, Symbol: "kr", ShowCodeNextToSymbol: true},
	"JEP": {Code: "JEP", Digits: 2, Symbol: "\u00a3", ShowCodeNextToSymbol: true},
	"JMD": {Code: "JMD", Digits: 2, Symbol: "J$", ShowCodeNextToSymbol: false},
	"JOD": {Code: "JOD", Digits: 3, Symbol: ".\u062f.\u0625", ShowCodeNextToSymbol: true},
	"JPY": {Code: "JPY", Digits: 0, Symbol: "\u00a5", ShowCodeNextToSymbol: false},
	"KES": {Code: "KES", Digits: 2, Symbol: "KSh", ShowCodeNextToSymbol: false},
	"KGS": {Code: "KGS", Digits: 2, Symbol: "\u0441\u043e\u043c", ShowCodeNextToSymbol: false},
	"KHR": {Code: "KHR", Digits: 2, Symbol: "\u17db", ShowCodeNextToSymbol: false},
	"KMF": {Code: "KMF", Digits: 0, Symbol: "CF", ShowCodeNextToSymbol: false},
	"KPW": {Code: "KPW", Digits: 0, Symbol: "\u20a9", ShowCodeNextToSymbol: true},
	"KRW": {Code: "KRW", Digits: 0, Symbol: "\u20a9", ShowCodeNextToSymbol: true},
	"KWD": {Code: "KWD", Digits: 3, Symbol: ".\u062f.\u0643", ShowCodeNextToSymbol: false},
	"KYD": {Code: "KYD", Digits: 2, Symbol: "CI$", ShowCodeNextToSymbol: false},
	"KZT": {Code: "KZT", Digits: 2, Symbol: "\u20b8", ShowCodeNextToSymbol: false},
	"LAK": {Code: "LAK", Digits: 2, Symbol: "\u20ad", ShowCodeNextToSymbol: false},
	"LBP": {Code: "LBP", Digits: 2, Symbol: "\u00a3", ShowCodeNextToSymbol: true},
	"LKR": {Code: "LKR", Digits: 2, Symbol: "\u20a8", ShowCodeNextToSymbol: true},
	"LRD": {Code: "LRD", Digits: 2, Symbol: "L$", ShowCodeNextToSymbol: false},
	"LSL": {Code: "LSL", Digits: 2, Symbol: "L", ShowCodeNextToSymbol: true},
	"LYD": {Code: "LYD", Digits: 3, Symbol: ".\u062f.\u0644", ShowCodeNextToSymbol: false},
	"MAD": {Code: "MAD", Digits: 2, Symbol: ".\u062f.\u0645", ShowCodeNextToSymbol: false},
	"MDL": {Code: "MDL", Digits: 2, Symbol: "lei", ShowCodeNextToSymbol: true},
	"MGA": {Code: "MGA", Digits: 0, Symbol: "Ar", ShowCodeNextToSymbol: false},
	"MKD": {Code: "MKD", Digits: 2, Symbol: "\u0434\u0435\u043d", ShowCodeNextToSymbol: false},
	"MMK": {Code: "MMK", Digits: 2, Symbol: "K", ShowCodeNextToSymbol: true},
	"MNT": {Code: "MNT", Digits: 2, Symbol: "\u20ae", ShowCodeNextToSymbol: false},
	"MOP": {Code: "MOP", Digits: 2, Symbol: "P", ShowCodeNextToSymbol: true},
	"MRO": {Code: "MRO", Digits: 0, Symbol: "UM", ShowCodeNextToSymbol: false},
	"MUR": {Code: "MUR", Digits: 2, Symbol: "\u20a8", ShowCodeNextToSymbol: true},
	"MVR": {Code: "MVR", Digits: 2, Symbol: "MVR", ShowCodeNextToSymbol: false},
	"MWK": {Code: "MWK", Digits: 2, Symbol: "MK", ShowCodeNextToSymbol: false},
	"MXN": {Code: "MXN", Digits: 2, Symbol: "Mex$", ShowCodeNextToSymbol: false},
	"MYR": {Code: "MYR", Digits: 2, Symbol: "RM", ShowCodeNextToSymbol: false},
	"MZN": {Code: "MZN", Digits: 2, Symbol: "MT", ShowCodeNextToSymbol: false},
	"NAD": {Code: "NAD", Digits: 2, Symbol: "N$", ShowCodeNextToSymbol: false},
	"NGN": {Code: "NGN", Digits: 2, Symbol: "\u20a6", ShowCodeNextToSymbol: false},
	"NIO": {Code: "NIO", Digits: 2, Symbol: "C$", ShowCodeNextToSymbol: false},
	"NOK": {Code: "NOK", Digits: 2, Symbol: "kr", ShowCodeNextToSymbol: true},
	"NPR": {Code: "NPR", Digits: 2, Symbol: "\u20a8", ShowCodeNextToSymbol: true},
	"NZD": {Code: "NZD", Digits: 2, Symbol: "NZ$", ShowCodeNextToSymbol: false},
	"OMR": {Code: "OMR", Digits: 3, Symbol: "\ufdfc", ShowCodeNextToSymbol: true},
	"PAB": {Code: "PAB", Digits: 2, Symbol: "B/.", ShowCodeNextToSymbol: false},
	"PEN": {Code: "PEN", Digits: 2, Symbol: "S/", ShowCodeNextToSymbol: false},
	"PGK": {Code: "PGK", Digits: 2, Symbol: "K", ShowCodeNextToSymbol: true},
	"PHP": {Code: "PHP", Digits: 2, Symbol: "\u20b1", ShowCodeNextToSymbol: false},
	"PKR": {Code: "PKR", Digits: 2, Symbol: "\u20a8", ShowCodeNextToSymbol: true},
	"PLN": {Code: "PLN", Digits: 2, Symbol: "z\u0142", ShowCodeNextToSymbol: false},
	"PYG": {Code: "PYG", Digits: 0, Symbol: "Gs", ShowCodeNextToSymbol: false},
	"QAR": {Code: "QAR", Digits: 2, Symbol: "\ufdfc", ShowCodeNextToSymbol: true},
	"RON": {Code: "RON", Digits: 2, Symbol: "lei", ShowCodeNextToSymbol: true},
	"RSD": {Code: "RSD", Digits: 2, Symbol: "\u0414\u0438\u043d.", ShowCodeNextToSymbol: false},
	"RUB": {Code: "RUB", Digits: 2, Symbol: "\u20bd", ShowCodeNextToSymbol: false},
	"RWF": {Code: "RWF", Digits: 0, Symbol: "FRw", ShowCodeNextToSymbol: false},
	"SAR": {Code: "SAR", Digits: 2, Symbol: "\ufdfc", ShowCodeNextToSymbol: true},
	"SBD": {Code: "SBD", Digits: 2, Symbol: "SI$", ShowCodeNextToSymbol: false},
	"SCR": {Code: "SCR", Digits: 2, Symbol: "\u20a8", ShowCodeNextToSymbol: true},
	"SDG": {Code: "SDG", Digits: 2, Symbol: "\u00a3", ShowCodeNextToSymbol: true},
	"SEK": {Code: "SEK", Digits: 2, Symbol: "kr", ShowCodeNextToSymbol: true},
	"SGD": {Code: "SGD", Digits: 2, Symbol: "S$", ShowCodeNextToSymbol: false},
	"SHP": {Code: "SHP", Digits: 2, Symbol: "\u00a3", ShowCodeNextToSymbol: true},
	"SLL": {Code: "SLL", Digits: 2, Symbol: "Le", ShowCodeNextToSymbol: false},
	"SOS": {Code: "SOS", Digits: 2, Symbol: "Sh", ShowCodeNextToSymbol: false},
	"SRD": {Code: "SRD", Digits: 2, Symbol: "", ShowCodeNextToSymbol: false},
	"SSP": {Code: "SSP", Digits: 2, Symbol: "\u00a3", ShowCodeNextToSymbol: true},
	"STD": {Code: "STD", Digits: 2, Symbol: "Db", ShowCodeNextToSymbol: false},
	"SVC": {Code: "SVC", Digits: 2, Symbol: "\u20a1", ShowCodeNextToSymbol: true},
	"SYP": {Code: "SYP", Digits: 2, Symbol: "\u00a3", ShowCodeNextToSymbol: true},
	"SZL": {Code: "SZL", Digits: 2, Symbol: "\u00a3", ShowCodeNextToSymbol: true},
	"THB": {Code: "THB", Digits: 2, Symbol: "\u0e3f", ShowCodeNextToSymbol: false},
	"TJS": {Code: "TJS", Digits: 2, Symbol: "SM", ShowCodeNextToSymbol: false},
	"TMT": {Code: "TMT", Digits: 2, Symbol: "T", ShowCodeNextToSymbol: true},
	"TND": {Code: "TND", Digits: 3, Symbol: ".\u062f.\u062a", ShowCodeNextToSymbol: false},
	"TOP": {Code: "TOP", Digits: 2, Symbol: "T$", ShowCodeNextToSymbol: false},
	"TRY": {Code: "TRY", Digits: 2, Symbol: "\u20ba", ShowCodeNextToSymbol: false},
	"TTD": {Code: "TTD", Digits: 2, Symbol: "TT$", ShowCodeNextToSymbol: false},
	"TWD": {Code: "TWD", Digits: 0, Symbol: "NT$", ShowCodeNextToSymbol: false},
	"TZS": {Code: "TZS", Digits: 0, Symbol: "TSh", ShowCodeNextToSymbol: false},
	"UAH": {Code: "UAH", Digits: 2, Symbol: "\u20b4", ShowCodeNextToSymbol: false},
	"UGX": {Code: "UGX", Digits: 0, Symbol: "USh", ShowCodeNextToSymbol: false},
	"USD": {Code: "USD", Digits: 2, Symbol: "$", ShowCodeNextToSymbol: false},
	"UYU": {Code: "UYU", Digits: 2, Symbol: "$U", ShowCodeNextToSymbol: false},
	"UZS": {Code: "UZS", Digits: 2, Symbol: "so\u2019m", ShowCodeNextToSymbol: false},
	"VES": {Code: "VES", Digits: 2, Symbol: "Bs.S", ShowCodeNextToSymbol: false},
	"VEF": {Code: "VEF", Digits: 2, Symbol: "Bs.F", ShowCodeNextToSymbol: false},
	"VND": {Code: "VND", Digits: 0, Symbol: "\u20ab", ShowCodeNextToSymbol: false},
	"VUV": {Code: "VUV", Digits: 0, Symbol: "Vt", ShowCodeNextToSymbol: false},
	"WST": {Code: "WST", Digits: 2, Symbol: "T", ShowCodeNextToSymbol: true},
	"XAF": {Code: "XAF", Digits: 0, Symbol: "Fr", ShowCodeNextToSymbol: true},
	"XOF": {Code: "XOF", Digits: 0, Symbol: "Fr", ShowCodeNextToSymbol: true},
	"XPF": {Code: "XPF", Digits: 0, Symbol: "Fr", ShowCodeNextToSymbol: true},
	"XCD": {Code: "XCD", Digits: 2, Symbol: "EC$", ShowCodeNextToSymbol: false},
	"YER": {Code: "YER", Digits: 2, Symbol: "\ufdfc", ShowCodeNextToSymbol: true},
	"ZAR": {Code: "ZAR", Digits: 2, Symbol: "R", ShowCodeNextToSymbol: false},
	"ZMW": {Code: "ZMW", Digits: 2, Symbol: "ZK", ShowCodeNextToSymbol: false},
	"ZWD": {Code: "ZWD", Digits: 2, Symbol: "Z$", ShowCodeNextToSymbol: false},
}
