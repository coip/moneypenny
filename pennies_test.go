package money

import "testing"

const Penny = Pennies(1)

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

func TestMoneyAddPennies(t *testing.T) {
	for _, c := range []struct {
		m      Money
		p      Pennies
		result string
	}{
		{Fromi(1), Pennies(1), "$0.02"},
		{Fromi(10), Pennies(10), "$0.20"},
		{Fromi(100), Pennies(100), "$2.00"},
		{Fromi(1000), Pennies(1000), "$20.00"},
		{Fromi(-1000), Pennies(-1000), "-$20.00"},
		{Fromf32(1.00), Pennies(100), "$2.00"},
		{Fromf32(10.5), Pennies(1050), "$21.00"},
		{Fromf32(10.01), Pennies(1001), "$20.02"},
		{Fromf32(10.00999999), Pennies(1001), "$20.02"},
		{Fromf64(1.00), Pennies(100), "$2.00"},
		{Fromf64(10.5), Pennies(1050), "$21.00"},
		{Fromf64(10.01), Pennies(1001), "$20.02"},
		{Fromf64(10.00999999), Pennies(1001), "$20.02"},
		{FromString("$10.01"), Pennies(1001), "$20.02"},
		{FromString("$10.1"), Pennies(1010), "$20.20"},
		{FromString("$10.009999999999999"), Pennies(1001), "$20.02"},
	} {
		got := c.m.AddP(c.p).String()
		if got != c.result {
			t.Errorf("(%v).AddP(%v).String() == %s, want %s", c.m, c.p, got, c.result)
		}
	}
}

func TestPenniesAddPennies(t *testing.T) {
	for _, c := range []struct {
		p1     Pennies
		p2     Pennies
		result Pennies
	}{
		{Pennies(1), Pennies(1), Pennies(2)},
		{Pennies(10), Pennies(10), Pennies(20)},
		{Pennies(100), Pennies(100), Pennies(200)},
		{Pennies(1000), Pennies(1000), Pennies(2000)},
		{Pennies(-1000), Pennies(-1000), Pennies(-2000)},
	} {
		got := c.p1.Add(c.p2)
		if got != c.result {
			t.Errorf("(%d).Add(%d) == %d, want %d", c.p1, c.p2, got, c.result)
		}
	}
}

func TestPenniesi64(t *testing.T) {
	if Pennies(1000).i64() != int64(1000) {
		t.Errorf("Error, pennies i64 underlying mishap.")
	}
}
func TestManyFloatPenniesCumulative(t *testing.T) {
	mm := int64(1000000)
	sum := Fromi(0)
	p := Fromf32(.01)
	for i := 0; i < int(mm); i++ {
		sum = sum.Add(p)
	}
	if !sum.Eq(Pennies(mm).ToMoney()) {
		t.Errorf("expected %d pennies, got %d", mm, sum.ToPennies())
	}
}

func BenchmarkAddingPennyToPennies(b *testing.B) {
	p := Penny * 10
	m := Pennies(0)
	for n := 0; n < b.N; n++ {
		m = m.Add(p)
	}
}

func BenchmarkAdding1MMPennyToPennies(b *testing.B) {
	p := Penny * 1000000
	m := Pennies(1000000)
	for n := 0; n < b.N; n++ {
		m = m.Add(p)
	}
}

//redundant. see ToMoney()
func BenchmarkAddingPenniesToMoney(b *testing.B) {
	p := Penny
	m := Fromi(0)
	for n := 0; n < b.N; n++ {
		m = m.Add(p.ToMoney())
	}
}

//redundant. see AddP
func BenchmarkAddingPennyToMoney(b *testing.B) {
	p := Penny
	m := Fromi(0)
	for n := 0; n < b.N; n++ {
		m = m.AddP(p)
	}
}

func BenchmarkAddingMoneyToPenniesAndPenny(b *testing.B) {
	p := Penny
	m := Fromi(100)
	for n := 0; n < b.N; n++ {
		p = p.Add(m.ToPennies())
	}
}

func BenchmarkCastingAddingPenniesNative(b *testing.B) {
	p, m := 1, 0

	for n := 0; n < b.N; n++ {
		m += p
	}
}

func BenchmarkCastingAddingPenniesNativeExplicit(b *testing.B) {
	p := int64(1)
	m := int64(0)

	for n := 0; n < b.N; n++ {
		m += p
	}
}

func BenchmarkCastingPenniesToInt64(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Penny.i64()
	}
}
