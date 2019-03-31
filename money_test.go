package money

import "testing"

func TestCovForTemplateFormatter(t *testing.T) {
	amt := 100

	c := struct {
		in   int
		m    Money
		want string
	}{
		amt,
		Fromi(amt),
		Fromi(amt).String(),
	}

	if FormatAsMoney(c.m) != c.want {
		t.Errorf("FormatAsMoney(Fromi(%d)) == %s, want %q", c.in, FormatAsMoney(c.m), c.want)
	}
}
func TestFromFromString(t *testing.T) {
	for _, c := range []struct {
		in, want string
	}{
		//dollars
		{"1000", "$1000.00"},
		{"100", "$100.00"},
		{"10", "$10.00"},
		//dollar formatted strings
		{"$1000", "$1000.00"},
		{"$100", "$100.00"},
		{"$10", "$10.00"},
		//negatives, also bank round and increments
		{"-10", "-$10.00"},
		{"-10.50", "-$10.50"},
		{"-10.4999", "-$10.50"},
		//negative formatted. not very useful case
		{"$-10", "-$10.00"},
		{"$-10.50", "-$10.50"},
		{"$-10.4999", "-$10.50"},
		//invalid
		{"coip", "$0.00"},
		//sporadic cases. possibly/likely missing some
		{"25.27", "$25.27"},
		{"0.33", "$0.33"},
		{"$584.56", "$584.56"},
		{"1", "$1.00"},
		{"00.2", "$0.20"},
		{"00.01", "$0.01"},
		{"-0.25", "-$0.25"},
		{"+0.85", "$0.85"},
		{"-10", "-$10.00"},
		{"-.0001", "-$0.00"},
		{"-999.999999999", "-$1000.00"},
	} {
		got := FromString(c.in).String()
		if got != c.want {
			t.Errorf("Froms(%q).String() == %q, want %q", c.in, got, c.want)
		}
	}
}

func TestFromFromInt(t *testing.T) {
	for _, c := range []struct {
		in   int
		want string
	}{
		{100000, "$1000.00"},
		{10000, "$100.00"},
		{1000, "$10.00"},
		{-1000, "-$10.00"},
	} {
		got := Fromi(c.in).String()
		if got != c.want {
			t.Errorf("FromInt(%d).String() == %q, want %q", c.in, got, c.want)
		}
	}
}

func TestFromFromFloat32(t *testing.T) {
	for _, c := range []struct {
		in   float32
		want string
	}{
		{1000, "$1000.00"},
		{100, "$100.00"},
		{10, "$10.00"},
		{-10, "-$10.00"},
		{1000.01, "$1000.01"},
		{100.99999999999999999, "$101.00"},
		{100.999, "$101.00"},
		{100.9901, "$100.99"},
		{100.456789, "$100.46"},
	} {
		got := Fromf32(c.in).String()
		if got != c.want {
			t.Errorf("Fromf32(%f).String() == %q, want %q", c.in, got, c.want)
		}
	}
}

func TestFromFromFloat64(t *testing.T) {
	for _, c := range []struct {
		in   float64
		want string
	}{
		{1000, "$1000.00"},
		{100, "$100.00"},
		{10, "$10.00"},
		{-10, "-$10.00"},
		{1000.01, "$1000.01"},
		{100.99999999999999999, "$101.00"},
		{100.999, "$101.00"},
		{100.9901, "$100.99"},
		{100.456789, "$100.46"},
	} {
		got := Fromf64(c.in).String()
		if got != c.want {
			t.Errorf("Fromf64(%f).String() == %q, want %q", c.in, got, c.want)
		}
	}
}

func TestFromFromInt64(t *testing.T) {
	for _, c := range []struct {
		in   int64
		want string
	}{
		{100000, "$1000.00"},
		{10000, "$100.00"},
		{1000, "$10.00"},
		{-1000, "-$10.00"},
	} {
		got := Fromi64(c.in).String()
		if got != c.want {
			t.Errorf("Fromi64(%d).String() == %q, want %q", c.in, got, c.want)
		}
	}
}

func BenchmarkFromString10(b *testing.B) {
	for n := 0; n < b.N; n++ {
		FromString("10")
	}
}
func BenchmarkFromString10DollarSign(b *testing.B) {
	for n := 0; n < b.N; n++ {
		FromString("$10")
	}
}

func BenchmarkFromString10DollarSignNeg(b *testing.B) {
	for n := 0; n < b.N; n++ {
		FromString("$-10")
	}
}

func BenchmarkFromString1000(b *testing.B) {
	for n := 0; n < b.N; n++ {
		FromString("1000")
	}
}

func BenchmarkFromString10Bucks(b *testing.B) {
	for n := 0; n < b.N; n++ {
		FromString("10.00")
	}
}

func BenchmarkFromStringInvalid(b *testing.B) {
	for n := 0; n < b.N; n++ {
		FromString("coip")
	}
}

func BenchmarkFromi100(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Fromi(100)
	}
}

func BenchmarkFromi1000(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Fromi(1000)
	}
}

func BenchmarkFromiMaxInt32(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Fromi(2147483647)
	}
}

func BenchmarkFromiMinInt32(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Fromi(-2147483648)
	}
}

func BenchmarkFromiMaxInt64(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Fromi64(9223372036854775807)
	}
}

func BenchmarkFromiMinInt64(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Fromi64(-9223372036854775808)
	}
}

//semi redundant. casting into int64 in the func.
func BenchmarkFromi10(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Fromi(10)
	}
}

//semi redundant.
func BenchmarkFromi64_10(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Fromi64(10)
	}
}

func BenchmarkFromf32_10(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Fromf32(10)
	}
}

func BenchmarkFromf64_10(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Fromf64(10)
	}
}

func BenchmarkString(b *testing.B) {
	m := Fromi(100)
	for n := 0; n < b.N; n++ {
		m.String()
	}
}

func BenchmarkFormatter(b *testing.B) {
	m := Fromi(100)
	for n := 0; n < b.N; n++ {
		FormatAsMoney(m)
	}
}

func BenchmarkStringNeg(b *testing.B) {
	m := Fromi(-100)
	for n := 0; n < b.N; n++ {
		m.String()
	}
}

func BenchmarkFormatterNeg(b *testing.B) {
	m := Fromi(-100)
	for n := 0; n < b.N; n++ {
		FormatAsMoney(m)
	}
}
