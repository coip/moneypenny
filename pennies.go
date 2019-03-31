package money

import "github.com/shopspring/decimal"

// Pennies provides plumbing, where Money is more porcelain.
type Pennies int64

func (p Pennies) ToMoney() Money {
	return Money(decimal.New(int64(p), -2))
}

func (p Pennies) i64() int64 {
	return int64(p)
}
