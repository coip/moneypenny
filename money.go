package money

import (
	"strings"

	"github.com/shopspring/decimal"
)

//Money bridges the gap (slightly) between semantics of client code and the plumbing its implemented via
type Money decimal.Decimal

//Doesnt handle:
// $-100
// $-100.00
// ($100)
// (100)

func FromString(m string) Money {
	if string(m[0]) == "$" {
		m = m[1:]
	}
	m = strings.Replace(m, ",", "", -1)
	tmp, e := decimal.NewFromString(m)
	if e != nil {
		return Money(decimal.NewFromFloat(0))
	}
	return Money(tmp)
}

func Fromf32(m float32) Money {
	return Money(decimal.NewFromFloat32(m))
}
func Fromf64(m float64) Money {
	return Money(decimal.NewFromFloat(m))
}
func Fromi(m int) Money {
	return Money(decimal.New(int64(m), -2))
}
func Fromi64(m int64) Money {
	return Money(decimal.New(m, -2))
}

func (m Money) String() string {
	if !decimal.Decimal(m).IsNegative() {
		return "$" + decimal.Decimal(m).StringFixedBank(2)
	}
	//else, format negative correctly.
	return "-$" + decimal.Decimal(m).Abs().StringFixedBank(2)
}

//FormatAsMoney is a callee/helper provided for template funcs
func FormatAsMoney(m Money) string {
	return m.String()
}
