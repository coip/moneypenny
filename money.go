package money

import (
	"github.com/shopspring/decimal"
)

//Money bridges the gap between semantics of client code and the plumbing its implemented via
type Money decimal.Decimal

func FromString(m string) Money {
	if string(m[0]) == "$" {
		m = m[1:]
	}
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

//Add enables m := m1 + m2
func (m1 Money) Add(m2 Money) Money {
	return Money(decimal.Decimal(m1).Add(decimal.Decimal(m2)))
}

//AddP enables m := Money + Pennies
func (m Money) AddP(p Pennies) Money {
	return m.Add(p.ToMoney())
}

//Eq enables if m1.Eq(m2) { log.Println("they're ==") }
func (m1 Money) Eq(m2 Money) bool {
	return decimal.Decimal(m1).Equal(decimal.Decimal(m2))
}

func (m Money) ToPennies() Pennies {
	return Pennies(decimal.Decimal(m).RoundBank(2).Shift(2).IntPart())
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
