package money

import "testing"

func TestToPennies(t *testing.T) {
	for _, c := range []struct {
		in   Money
		want Pennies
	}{
		{Fromi(1), Pennies(1)},
		{Fromi(10), Pennies(10)},
		{Fromi(100), Pennies(100)},
		{Fromi(1000), Pennies(1000)},
		{Fromi(-1000), Pennies(-1000)},
		{Fromf32(1.00), Pennies(100)},
		{Fromf32(10.5), Pennies(1050)},
		{Fromf32(10.01), Pennies(1001)},
		{Fromf32(10.00999999), Pennies(1001)},
		{Fromf64(1.00), Pennies(100)},
		{Fromf64(10.5), Pennies(1050)},
		{Fromf64(10.01), Pennies(1001)},
		{Fromf64(10.00999999), Pennies(1001)},
		{FromString("$10.01"), Pennies(1001)},
		{FromString("$10.1"), Pennies(1010)},
		{FromString("$10.009999999999999"), Pennies(1001)},
	} {
		got := c.in.ToPennies()
		if got != c.want {
			t.Errorf("Fromi(%s).ToPennies() == %d, want %d", c.in, got, c.want)
		}
	}
}

func TestManyFloatPenniesCumulative(t *testing.T) {
	sum := Fromi(0)
	p := Fromf32(.01)
	for i := 0; i < 1000000; i++ {
		sum = sum.Add(p)
	}
	if !sum.Eq(Pennies(1000000).ToMoney()) {
		t.Errorf("expected 1000000 pennies, got %d", sum.ToPennies())
	}
}
