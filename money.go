package money

import (
	"errors"
	"fmt"
	"math"
	"strconv"

	"database/sql/driver"
)

type Amount int64

func (a Amount) Int64() int64 {
	return int64(a)
}

type Money struct {
	Amount   Amount   `json:"amount"`
	Currency Currency `json:"currency"`
}

func (m Money) Int64() int64 {
	return int64(m.Amount)
}

// Forge
// amount   int64   A positive integer in cents
// currCode string Three-letter ISO currency code, in lowercase
func Forge(amount int64, currCode string) (m Money, err error) {
	c, err := CurrencyByISOCode(currCode)
	if err != nil {
		return
	}
	m = ForgeWithCurrency(amount, c)
	return
}

// Forge
// amount   float64 A positive float64 in cents
// currCode string  Three-letter ISO currency code, in lowercase
func ForgeFloat(amount float64, currCode string) (m Money, err error) {
	c, err := CurrencyByISOCode(currCode)
	if err != nil {
		return m, err
	}

	m = ForgeFloatWithCurrency(amount, c)
	return m, err
}

func MustForgeFloat(amountFloat float64, currCode string) Money {
	m, err := Forge(0, currCode)
	if err != nil {
		panic(err)
	}

	return ForgeFloatWithCurrency(amountFloat, m.Currency)
}

func ForgeFloatWithCurrency(amountFloat float64, c Currency) Money {
	fd := float64(c.GetCents())
	amount := int64(math.Round(fd * amountFloat))

	return ForgeWithCurrency(amount, c)
}

// MustForge Forge or panic
// amount   int    A positive integer in cents
// currCode string Three-letter ISO currency code, in lowercase
func MustForge(amount int64, currCode string) Money {
	m, err := Forge(amount, currCode)
	if err != nil {
		panic(err)
	}

	return m
}

// ForgeWithCurrency
// amount   int      A positive integer in cents
// currency Currency The currency Value Object
func ForgeWithCurrency(amount int64, c Currency) Money {
	var err error
	if !c.IsValid() {
		c, err = CurrencyByISOCode(DefaultCurrencyCode)
		if err != nil {
			panic(err)
		}
	}

	return Money{Amount: Amount(amount), Currency: c}
}

func (m Money) IsZero() bool {
	return m.Amount == 0
}

func (m Money) Float() float64 {
	if m.IsZero() {
		return 0
	}
	d := m.DigitsAsCents()
	return float64(m.Amount) / float64(d)
}

func (m Money) DigitsAsCents() int {
	return m.Currency.GetCents()
}

func (m Money) IsEquals(cmp Money) bool {
	return m.Amount == cmp.Amount && m.Currency.IsEquals(cmp.Currency)
}

func (m Money) PercentOff(perc int) Money {
	div := float64(perc) / 100
	return ForgeFloatWithCurrency(m.Float()*div, m.Currency)
}

func (m Money) Add(addendum Money) (s Money, err error) {
	if !m.Currency.IsEquals(addendum.Currency) {
		return s, errors.New(fmt.Sprint("Can't compare or use math with different currency", m.Currency, s.Currency))
	}

	return ForgeWithCurrency(addendum.Amount.Int64()+m.Amount.Int64(), m.Currency), err
}

func (m Money) Subtract(subtrahend Money) (s Money, err error) {
	if !m.Currency.IsEquals(subtrahend.Currency) {
		return s, errors.New(fmt.Sprint("Can't compare or use math with different currency", m.Currency, s.Currency))
	}

	return ForgeWithCurrency(m.Amount.Int64()-subtrahend.Amount.Int64(), m.Currency), err
}

func (m Money) SplitAmountAndCents() (i int64, cents int, err error) {
	f := m.Float()
	digitsCount := m.Currency.MinorUnit

	i = int64(f)
	if f-float64(i) <= 0 {
		return i, cents, err
	}

	cutAt := 2 + digitsCount
	ct := fmt.Sprintf("%d", cutAt)
	decPartStr := fmt.Sprintf("%."+ct+"g", f-float64(i))

	if len(decPartStr) <= 2 {
		return i, cents, errors.New("can't convert to decimal " + decPartStr)
	}

	if len(decPartStr) > cutAt {
		cents, err = strconv.Atoi(decPartStr[2:cutAt])
		return i, cents, err
	}

	cents, err = strconv.Atoi(decPartStr[2:])
	return i, cents, err
}

func Scan(value interface{}, curr string) (m Money, err error) {
	if v, ok := value.(int64); ok {
		mm, err := Forge(v, curr)
		return mm, err
	}
	return m, fmt.Errorf("impossible to get int64 the value from %v", value)
}

func (m *Money) ScanInt64(value interface{}) error {
	if value == nil {
		return nil
	}
	code := DefaultCurrencyCode
	if m.Currency.IsValid() {
		code = m.Currency.Code
	}
	mm, err := Scan(value, code)
	if err != nil {
		return err
	}
	*m = mm
	return nil
}

func (m *Money) Scan(value interface{}) error {
	if value == nil {
		return nil
	}

	if v, ok := value.(string); ok {
		mm, err := ParseWithFallback(v, m.Currency)
		if err != nil {
			return err
		}
		*m = mm
		return err
	}

	if v, ok := value.(int64); ok {
		mm := ForgeWithCurrency(v, m.Currency)
		err := mm.ScanInt64(value)
		if err != nil {
			return err
		}
		*m = mm
		return err
	}

	if v, ok := value.(float64); ok {
		code := DefaultCurrencyCode
		if m.Currency.IsValid() {
			code = m.Currency.Code
		}
		mm, err := ForgeFloat(v, code)
		if err != nil {
			return err
		}
		*m = mm
		return nil
	}

	return fmt.Errorf("can't convert given %v", value)
}

// Value implements the driver Valuer interface.
func (m *Money) Value() (driver.Value, error) {
	return m.Int64(), nil
}

// Value implements the driver Valuer interface.
func (m *Money) String() string {
	return fmt.Sprintf("%s %d", m.Currency.String(), m.Int64())
}

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

func AED(i int64) Money { return MustForge(i, "AED") }
func AFN(i int64) Money { return MustForge(i, "AFN") }
func ALL(i int64) Money { return MustForge(i, "ALL") }
func AMD(i int64) Money { return MustForge(i, "AMD") }
func ANG(i int64) Money { return MustForge(i, "ANG") }
func AOA(i int64) Money { return MustForge(i, "AOA") }
func ARS(i int64) Money { return MustForge(i, "ARS") }
func AUD(i int64) Money { return MustForge(i, "AUD") }
func AWG(i int64) Money { return MustForge(i, "AWG") }
func AZN(i int64) Money { return MustForge(i, "AZN") }
func BAM(i int64) Money { return MustForge(i, "BAM") }
func BBD(i int64) Money { return MustForge(i, "BBD") }
func BDT(i int64) Money { return MustForge(i, "BDT") }
func BGN(i int64) Money { return MustForge(i, "BGN") }
func BHD(i int64) Money { return MustForge(i, "BHD") }
func BIF(i int64) Money { return MustForge(i, "BIF") }
func BMD(i int64) Money { return MustForge(i, "BMD") }
func BND(i int64) Money { return MustForge(i, "BND") }
func BOB(i int64) Money { return MustForge(i, "BOB") }
func BRL(i int64) Money { return MustForge(i, "BRL") }
func BSD(i int64) Money { return MustForge(i, "BSD") }
func BTN(i int64) Money { return MustForge(i, "BTN") }
func BWP(i int64) Money { return MustForge(i, "BWP") }
func BYN(i int64) Money { return MustForge(i, "BYN") }
func BZD(i int64) Money { return MustForge(i, "BZD") }
func CAD(i int64) Money { return MustForge(i, "CAD") }
func CDF(i int64) Money { return MustForge(i, "CDF") }
func CHF(i int64) Money { return MustForge(i, "CHF") }
func CLF(i int64) Money { return MustForge(i, "CLF") }
func CLP(i int64) Money { return MustForge(i, "CLP") }
func CNY(i int64) Money { return MustForge(i, "CNY") }
func COP(i int64) Money { return MustForge(i, "COP") }
func CRC(i int64) Money { return MustForge(i, "CRC") }
func CUC(i int64) Money { return MustForge(i, "CUC") }
func CUP(i int64) Money { return MustForge(i, "CUP") }
func CVE(i int64) Money { return MustForge(i, "CVE") }
func CZK(i int64) Money { return MustForge(i, "CZK") }
func DJF(i int64) Money { return MustForge(i, "DJF") }
func DKK(i int64) Money { return MustForge(i, "DKK") }
func DOP(i int64) Money { return MustForge(i, "DOP") }
func DZD(i int64) Money { return MustForge(i, "DZD") }
func EGP(i int64) Money { return MustForge(i, "EGP") }
func ERN(i int64) Money { return MustForge(i, "ERN") }
func ETB(i int64) Money { return MustForge(i, "ETB") }
func EUR(i int64) Money { return MustForge(i, "EUR") }
func FJD(i int64) Money { return MustForge(i, "FJD") }
func FKP(i int64) Money { return MustForge(i, "FKP") }
func GBP(i int64) Money { return MustForge(i, "GBP") }
func GEL(i int64) Money { return MustForge(i, "GEL") }
func GGP(i int64) Money { return MustForge(i, "GGP") }
func GHS(i int64) Money { return MustForge(i, "GHS") }
func GIP(i int64) Money { return MustForge(i, "GIP") }
func GMD(i int64) Money { return MustForge(i, "GMD") }
func GNF(i int64) Money { return MustForge(i, "GNF") }
func GTQ(i int64) Money { return MustForge(i, "GTQ") }
func GYD(i int64) Money { return MustForge(i, "GYD") }
func HKD(i int64) Money { return MustForge(i, "HKD") }
func HNL(i int64) Money { return MustForge(i, "HNL") }
func HRK(i int64) Money { return MustForge(i, "HRK") }
func HTG(i int64) Money { return MustForge(i, "HTG") }
func HUF(i int64) Money { return MustForge(i, "HUF") }
func IDR(i int64) Money { return MustForge(i, "IDR") }
func ILS(i int64) Money { return MustForge(i, "ILS") }
func IMP(i int64) Money { return MustForge(i, "IMP") }
func INR(i int64) Money { return MustForge(i, "INR") }
func IQD(i int64) Money { return MustForge(i, "IQD") }
func IRR(i int64) Money { return MustForge(i, "IRR") }
func ISK(i int64) Money { return MustForge(i, "ISK") }
func JEP(i int64) Money { return MustForge(i, "JEP") }
func JMD(i int64) Money { return MustForge(i, "JMD") }
func JOD(i int64) Money { return MustForge(i, "JOD") }
func JPY(i int64) Money { return MustForge(i, "JPY") }
func KES(i int64) Money { return MustForge(i, "KES") }
func KGS(i int64) Money { return MustForge(i, "KGS") }
func KHR(i int64) Money { return MustForge(i, "KHR") }
func KMF(i int64) Money { return MustForge(i, "KMF") }
func KPW(i int64) Money { return MustForge(i, "KPW") }
func KRW(i int64) Money { return MustForge(i, "KRW") }
func KWD(i int64) Money { return MustForge(i, "KWD") }
func KYD(i int64) Money { return MustForge(i, "KYD") }
func KZT(i int64) Money { return MustForge(i, "KZT") }
func LAK(i int64) Money { return MustForge(i, "LAK") }
func LBP(i int64) Money { return MustForge(i, "LBP") }
func LKR(i int64) Money { return MustForge(i, "LKR") }
func LRD(i int64) Money { return MustForge(i, "LRD") }
func LSL(i int64) Money { return MustForge(i, "LSL") }
func LYD(i int64) Money { return MustForge(i, "LYD") }
func MAD(i int64) Money { return MustForge(i, "MAD") }
func MDL(i int64) Money { return MustForge(i, "MDL") }
func MGA(i int64) Money { return MustForge(i, "MGA") }
func MKD(i int64) Money { return MustForge(i, "MKD") }
func MMK(i int64) Money { return MustForge(i, "MMK") }
func MNT(i int64) Money { return MustForge(i, "MNT") }
func MOP(i int64) Money { return MustForge(i, "MOP") }
func MRO(i int64) Money { return MustForge(i, "MRO") }
func MUR(i int64) Money { return MustForge(i, "MUR") }
func MVR(i int64) Money { return MustForge(i, "MVR") }
func MWK(i int64) Money { return MustForge(i, "MWK") }
func MXN(i int64) Money { return MustForge(i, "MXN") }
func MYR(i int64) Money { return MustForge(i, "MYR") }
func MZN(i int64) Money { return MustForge(i, "MZN") }
func NAD(i int64) Money { return MustForge(i, "NAD") }
func NGN(i int64) Money { return MustForge(i, "NGN") }
func NIO(i int64) Money { return MustForge(i, "NIO") }
func NOK(i int64) Money { return MustForge(i, "NOK") }
func NPR(i int64) Money { return MustForge(i, "NPR") }
func NZD(i int64) Money { return MustForge(i, "NZD") }
func OMR(i int64) Money { return MustForge(i, "OMR") }
func PAB(i int64) Money { return MustForge(i, "PAB") }
func PEN(i int64) Money { return MustForge(i, "PEN") }
func PGK(i int64) Money { return MustForge(i, "PGK") }
func PHP(i int64) Money { return MustForge(i, "PHP") }
func PKR(i int64) Money { return MustForge(i, "PKR") }
func PLN(i int64) Money { return MustForge(i, "PLN") }
func PYG(i int64) Money { return MustForge(i, "PYG") }
func QAR(i int64) Money { return MustForge(i, "QAR") }
func RON(i int64) Money { return MustForge(i, "RON") }
func RSD(i int64) Money { return MustForge(i, "RSD") }
func RUB(i int64) Money { return MustForge(i, "RUB") }
func RWF(i int64) Money { return MustForge(i, "RWF") }
func SAR(i int64) Money { return MustForge(i, "SAR") }
func SBD(i int64) Money { return MustForge(i, "SBD") }
func SCR(i int64) Money { return MustForge(i, "SCR") }
func SDG(i int64) Money { return MustForge(i, "SDG") }
func SEK(i int64) Money { return MustForge(i, "SEK") }
func SGD(i int64) Money { return MustForge(i, "SGD") }
func SHP(i int64) Money { return MustForge(i, "SHP") }
func SLL(i int64) Money { return MustForge(i, "SLL") }
func SOS(i int64) Money { return MustForge(i, "SOS") }
func SRD(i int64) Money { return MustForge(i, "SRD") }
func SSP(i int64) Money { return MustForge(i, "SSP") }
func STD(i int64) Money { return MustForge(i, "STD") }
func SVC(i int64) Money { return MustForge(i, "SVC") }
func SYP(i int64) Money { return MustForge(i, "SYP") }
func SZL(i int64) Money { return MustForge(i, "SZL") }
func THB(i int64) Money { return MustForge(i, "THB") }
func TJS(i int64) Money { return MustForge(i, "TJS") }
func TMT(i int64) Money { return MustForge(i, "TMT") }
func TND(i int64) Money { return MustForge(i, "TND") }
func TOP(i int64) Money { return MustForge(i, "TOP") }
func TRY(i int64) Money { return MustForge(i, "TRY") }
func TTD(i int64) Money { return MustForge(i, "TTD") }
func TWD(i int64) Money { return MustForge(i, "TWD") }
func TZS(i int64) Money { return MustForge(i, "TZS") }
func UAH(i int64) Money { return MustForge(i, "UAH") }
func UGX(i int64) Money { return MustForge(i, "UGX") }
func USD(i int64) Money { return MustForge(i, "USD") }
func UYU(i int64) Money { return MustForge(i, "UYU") }
func UZS(i int64) Money { return MustForge(i, "UZS") }
func VES(i int64) Money { return MustForge(i, "VES") }
func VEF(i int64) Money { return MustForge(i, "VEF") }
func VND(i int64) Money { return MustForge(i, "VND") }
func VUV(i int64) Money { return MustForge(i, "VUV") }
func WST(i int64) Money { return MustForge(i, "WST") }
func XAF(i int64) Money { return MustForge(i, "XAF") }
func XOF(i int64) Money { return MustForge(i, "XOF") }
func XPF(i int64) Money { return MustForge(i, "XPF") }
func XCD(i int64) Money { return MustForge(i, "XCD") }
func YER(i int64) Money { return MustForge(i, "YER") }
func ZAR(i int64) Money { return MustForge(i, "ZAR") }
func ZMW(i int64) Money { return MustForge(i, "ZMW") }
func ZWD(i int64) Money { return MustForge(i, "ZWD") }

func FloatAED(i float64) Money { return MustForgeFloat(i, "AED") }
func FloatAFN(i float64) Money { return MustForgeFloat(i, "AFN") }
func FloatALL(i float64) Money { return MustForgeFloat(i, "ALL") }
func FloatAMD(i float64) Money { return MustForgeFloat(i, "AMD") }
func FloatANG(i float64) Money { return MustForgeFloat(i, "ANG") }
func FloatAOA(i float64) Money { return MustForgeFloat(i, "AOA") }
func FloatARS(i float64) Money { return MustForgeFloat(i, "ARS") }
func FloatAUD(i float64) Money { return MustForgeFloat(i, "AUD") }
func FloatAWG(i float64) Money { return MustForgeFloat(i, "AWG") }
func FloatAZN(i float64) Money { return MustForgeFloat(i, "AZN") }
func FloatBAM(i float64) Money { return MustForgeFloat(i, "BAM") }
func FloatBBD(i float64) Money { return MustForgeFloat(i, "BBD") }
func FloatBDT(i float64) Money { return MustForgeFloat(i, "BDT") }
func FloatBGN(i float64) Money { return MustForgeFloat(i, "BGN") }
func FloatBHD(i float64) Money { return MustForgeFloat(i, "BHD") }
func FloatBIF(i float64) Money { return MustForgeFloat(i, "BIF") }
func FloatBMD(i float64) Money { return MustForgeFloat(i, "BMD") }
func FloatBND(i float64) Money { return MustForgeFloat(i, "BND") }
func FloatBOB(i float64) Money { return MustForgeFloat(i, "BOB") }
func FloatBRL(i float64) Money { return MustForgeFloat(i, "BRL") }
func FloatBSD(i float64) Money { return MustForgeFloat(i, "BSD") }
func FloatBTN(i float64) Money { return MustForgeFloat(i, "BTN") }
func FloatBWP(i float64) Money { return MustForgeFloat(i, "BWP") }
func FloatBYN(i float64) Money { return MustForgeFloat(i, "BYN") }
func FloatBZD(i float64) Money { return MustForgeFloat(i, "BZD") }
func FloatCAD(i float64) Money { return MustForgeFloat(i, "CAD") }
func FloatCDF(i float64) Money { return MustForgeFloat(i, "CDF") }
func FloatCHF(i float64) Money { return MustForgeFloat(i, "CHF") }
func FloatCLF(i float64) Money { return MustForgeFloat(i, "CLF") }
func FloatCLP(i float64) Money { return MustForgeFloat(i, "CLP") }
func FloatCNY(i float64) Money { return MustForgeFloat(i, "CNY") }
func FloatCOP(i float64) Money { return MustForgeFloat(i, "COP") }
func FloatCRC(i float64) Money { return MustForgeFloat(i, "CRC") }
func FloatCUC(i float64) Money { return MustForgeFloat(i, "CUC") }
func FloatCUP(i float64) Money { return MustForgeFloat(i, "CUP") }
func FloatCVE(i float64) Money { return MustForgeFloat(i, "CVE") }
func FloatCZK(i float64) Money { return MustForgeFloat(i, "CZK") }
func FloatDJF(i float64) Money { return MustForgeFloat(i, "DJF") }
func FloatDKK(i float64) Money { return MustForgeFloat(i, "DKK") }
func FloatDOP(i float64) Money { return MustForgeFloat(i, "DOP") }
func FloatDZD(i float64) Money { return MustForgeFloat(i, "DZD") }
func FloatEGP(i float64) Money { return MustForgeFloat(i, "EGP") }
func FloatERN(i float64) Money { return MustForgeFloat(i, "ERN") }
func FloatETB(i float64) Money { return MustForgeFloat(i, "ETB") }
func FloatEUR(i float64) Money { return MustForgeFloat(i, "EUR") }
func FloatFJD(i float64) Money { return MustForgeFloat(i, "FJD") }
func FloatFKP(i float64) Money { return MustForgeFloat(i, "FKP") }
func FloatGBP(i float64) Money { return MustForgeFloat(i, "GBP") }
func FloatGEL(i float64) Money { return MustForgeFloat(i, "GEL") }
func FloatGGP(i float64) Money { return MustForgeFloat(i, "GGP") }
func FloatGHS(i float64) Money { return MustForgeFloat(i, "GHS") }
func FloatGIP(i float64) Money { return MustForgeFloat(i, "GIP") }
func FloatGMD(i float64) Money { return MustForgeFloat(i, "GMD") }
func FloatGNF(i float64) Money { return MustForgeFloat(i, "GNF") }
func FloatGTQ(i float64) Money { return MustForgeFloat(i, "GTQ") }
func FloatGYD(i float64) Money { return MustForgeFloat(i, "GYD") }
func FloatHKD(i float64) Money { return MustForgeFloat(i, "HKD") }
func FloatHNL(i float64) Money { return MustForgeFloat(i, "HNL") }
func FloatHRK(i float64) Money { return MustForgeFloat(i, "HRK") }
func FloatHTG(i float64) Money { return MustForgeFloat(i, "HTG") }
func FloatHUF(i float64) Money { return MustForgeFloat(i, "HUF") }
func FloatIDR(i float64) Money { return MustForgeFloat(i, "IDR") }
func FloatILS(i float64) Money { return MustForgeFloat(i, "ILS") }
func FloatIMP(i float64) Money { return MustForgeFloat(i, "IMP") }
func FloatINR(i float64) Money { return MustForgeFloat(i, "INR") }
func FloatIQD(i float64) Money { return MustForgeFloat(i, "IQD") }
func FloatIRR(i float64) Money { return MustForgeFloat(i, "IRR") }
func FloatISK(i float64) Money { return MustForgeFloat(i, "ISK") }
func FloatJEP(i float64) Money { return MustForgeFloat(i, "JEP") }
func FloatJMD(i float64) Money { return MustForgeFloat(i, "JMD") }
func FloatJOD(i float64) Money { return MustForgeFloat(i, "JOD") }
func FloatJPY(i float64) Money { return MustForgeFloat(i, "JPY") }
func FloatKES(i float64) Money { return MustForgeFloat(i, "KES") }
func FloatKGS(i float64) Money { return MustForgeFloat(i, "KGS") }
func FloatKHR(i float64) Money { return MustForgeFloat(i, "KHR") }
func FloatKMF(i float64) Money { return MustForgeFloat(i, "KMF") }
func FloatKPW(i float64) Money { return MustForgeFloat(i, "KPW") }
func FloatKRW(i float64) Money { return MustForgeFloat(i, "KRW") }
func FloatKWD(i float64) Money { return MustForgeFloat(i, "KWD") }
func FloatKYD(i float64) Money { return MustForgeFloat(i, "KYD") }
func FloatKZT(i float64) Money { return MustForgeFloat(i, "KZT") }
func FloatLAK(i float64) Money { return MustForgeFloat(i, "LAK") }
func FloatLBP(i float64) Money { return MustForgeFloat(i, "LBP") }
func FloatLKR(i float64) Money { return MustForgeFloat(i, "LKR") }
func FloatLRD(i float64) Money { return MustForgeFloat(i, "LRD") }
func FloatLSL(i float64) Money { return MustForgeFloat(i, "LSL") }
func FloatLYD(i float64) Money { return MustForgeFloat(i, "LYD") }
func FloatMAD(i float64) Money { return MustForgeFloat(i, "MAD") }
func FloatMDL(i float64) Money { return MustForgeFloat(i, "MDL") }
func FloatMGA(i float64) Money { return MustForgeFloat(i, "MGA") }
func FloatMKD(i float64) Money { return MustForgeFloat(i, "MKD") }
func FloatMMK(i float64) Money { return MustForgeFloat(i, "MMK") }
func FloatMNT(i float64) Money { return MustForgeFloat(i, "MNT") }
func FloatMOP(i float64) Money { return MustForgeFloat(i, "MOP") }
func FloatMRO(i float64) Money { return MustForgeFloat(i, "MRO") }
func FloatMUR(i float64) Money { return MustForgeFloat(i, "MUR") }
func FloatMVR(i float64) Money { return MustForgeFloat(i, "MVR") }
func FloatMWK(i float64) Money { return MustForgeFloat(i, "MWK") }
func FloatMXN(i float64) Money { return MustForgeFloat(i, "MXN") }
func FloatMYR(i float64) Money { return MustForgeFloat(i, "MYR") }
func FloatMZN(i float64) Money { return MustForgeFloat(i, "MZN") }
func FloatNAD(i float64) Money { return MustForgeFloat(i, "NAD") }
func FloatNGN(i float64) Money { return MustForgeFloat(i, "NGN") }
func FloatNIO(i float64) Money { return MustForgeFloat(i, "NIO") }
func FloatNOK(i float64) Money { return MustForgeFloat(i, "NOK") }
func FloatNPR(i float64) Money { return MustForgeFloat(i, "NPR") }
func FloatNZD(i float64) Money { return MustForgeFloat(i, "NZD") }
func FloatOMR(i float64) Money { return MustForgeFloat(i, "OMR") }
func FloatPAB(i float64) Money { return MustForgeFloat(i, "PAB") }
func FloatPEN(i float64) Money { return MustForgeFloat(i, "PEN") }
func FloatPGK(i float64) Money { return MustForgeFloat(i, "PGK") }
func FloatPHP(i float64) Money { return MustForgeFloat(i, "PHP") }
func FloatPKR(i float64) Money { return MustForgeFloat(i, "PKR") }
func FloatPLN(i float64) Money { return MustForgeFloat(i, "PLN") }
func FloatPYG(i float64) Money { return MustForgeFloat(i, "PYG") }
func FloatQAR(i float64) Money { return MustForgeFloat(i, "QAR") }
func FloatRON(i float64) Money { return MustForgeFloat(i, "RON") }
func FloatRSD(i float64) Money { return MustForgeFloat(i, "RSD") }
func FloatRUB(i float64) Money { return MustForgeFloat(i, "RUB") }
func FloatRWF(i float64) Money { return MustForgeFloat(i, "RWF") }
func FloatSAR(i float64) Money { return MustForgeFloat(i, "SAR") }
func FloatSBD(i float64) Money { return MustForgeFloat(i, "SBD") }
func FloatSCR(i float64) Money { return MustForgeFloat(i, "SCR") }
func FloatSDG(i float64) Money { return MustForgeFloat(i, "SDG") }
func FloatSEK(i float64) Money { return MustForgeFloat(i, "SEK") }
func FloatSGD(i float64) Money { return MustForgeFloat(i, "SGD") }
func FloatSHP(i float64) Money { return MustForgeFloat(i, "SHP") }
func FloatSLL(i float64) Money { return MustForgeFloat(i, "SLL") }
func FloatSOS(i float64) Money { return MustForgeFloat(i, "SOS") }
func FloatSRD(i float64) Money { return MustForgeFloat(i, "SRD") }
func FloatSSP(i float64) Money { return MustForgeFloat(i, "SSP") }
func FloatSTD(i float64) Money { return MustForgeFloat(i, "STD") }
func FloatSVC(i float64) Money { return MustForgeFloat(i, "SVC") }
func FloatSYP(i float64) Money { return MustForgeFloat(i, "SYP") }
func FloatSZL(i float64) Money { return MustForgeFloat(i, "SZL") }
func FloatTHB(i float64) Money { return MustForgeFloat(i, "THB") }
func FloatTJS(i float64) Money { return MustForgeFloat(i, "TJS") }
func FloatTMT(i float64) Money { return MustForgeFloat(i, "TMT") }
func FloatTND(i float64) Money { return MustForgeFloat(i, "TND") }
func FloatTOP(i float64) Money { return MustForgeFloat(i, "TOP") }
func FloatTRY(i float64) Money { return MustForgeFloat(i, "TRY") }
func FloatTTD(i float64) Money { return MustForgeFloat(i, "TTD") }
func FloatTWD(i float64) Money { return MustForgeFloat(i, "TWD") }
func FloatTZS(i float64) Money { return MustForgeFloat(i, "TZS") }
func FloatUAH(i float64) Money { return MustForgeFloat(i, "UAH") }
func FloatUGX(i float64) Money { return MustForgeFloat(i, "UGX") }
func FloatUSD(i float64) Money { return MustForgeFloat(i, "USD") }
func FloatUYU(i float64) Money { return MustForgeFloat(i, "UYU") }
func FloatUZS(i float64) Money { return MustForgeFloat(i, "UZS") }
func FloatVES(i float64) Money { return MustForgeFloat(i, "VES") }
func FloatVEF(i float64) Money { return MustForgeFloat(i, "VEF") }
func FloatVND(i float64) Money { return MustForgeFloat(i, "VND") }
func FloatVUV(i float64) Money { return MustForgeFloat(i, "VUV") }
func FloatWST(i float64) Money { return MustForgeFloat(i, "WST") }
func FloatXAF(i float64) Money { return MustForgeFloat(i, "XAF") }
func FloatXOF(i float64) Money { return MustForgeFloat(i, "XOF") }
func FloatXPF(i float64) Money { return MustForgeFloat(i, "XPF") }
func FloatXCD(i float64) Money { return MustForgeFloat(i, "XCD") }
func FloatYER(i float64) Money { return MustForgeFloat(i, "YER") }
func FloatZAR(i float64) Money { return MustForgeFloat(i, "ZAR") }
func FloatZMW(i float64) Money { return MustForgeFloat(i, "ZMW") }
func FloatZWD(i float64) Money { return MustForgeFloat(i, "ZWD") }
