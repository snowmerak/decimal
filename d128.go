package decimal

import "fmt"

type D128 struct {
	sign        Sign
	numerator   uint64
	denominator uint64
}

func New128(sign Sign, numerator, denominator uint64) D128 {
	gcd := Gcd64(numerator, denominator)
	numerator /= gcd
	denominator /= gcd
	return D128{
		sign:        sign,
		numerator:   numerator,
		denominator: denominator,
	}
}

func (a D128) String() string {
	sign := ""
	if a.sign == Negative {
		sign = "-"
	}
	return fmt.Sprintf("%s%d/%d", sign, a.numerator, a.denominator)
}

func (a D128) Add(b D128) D128 {
	if a.sign != b.sign {
		if a.sign == Negative {
			a.sign = Positive
			return b.Sub(a)
		}
		b.sign = Positive
		return a.Sub(b)
	}
	gcd := Gcd64(a.denominator, b.denominator)
	a.denominator /= gcd
	b.denominator /= gcd
	a.numerator *= b.denominator
	b.numerator *= a.denominator
	denominator := a.denominator * b.denominator * gcd
	numerator := a.numerator + b.numerator
	gcd = Gcd64(numerator, denominator)
	numerator /= gcd
	denominator /= gcd
	return D128{
		sign:        a.sign,
		denominator: denominator,
		numerator:   numerator,
	}
}

func (a D128) Sub(b D128) D128 {
	if a.sign != b.sign {
		if a.sign == Positive {
			a.sign = Negative
			return a.Add(b)
		}
		b.sign = Negative
		return b.Add(a)
	}
	gcd := Gcd64(a.denominator, b.denominator)
	a.denominator /= gcd
	b.denominator /= gcd
	a.numerator *= b.denominator
	b.numerator *= a.denominator
	denominator := a.denominator * b.denominator * gcd
	numerator := a.numerator - b.numerator
	gcd = Gcd64(numerator, denominator)
	numerator /= gcd
	denominator /= gcd
	return D128{
		sign:        a.sign,
		denominator: denominator,
		numerator:   numerator,
	}
}

func (a D128) Mul(b D128) D128 {
	sign := Positive
	if a.sign != b.sign {
		sign = Negative
	}
	gcd1 := Gcd64(a.numerator, b.denominator)
	gcd2 := Gcd64(b.numerator, a.denominator)
	numerator := a.numerator / gcd1 * (b.numerator / gcd2)
	denominator := a.denominator / gcd2 * (b.denominator / gcd1)
	gcd := Gcd64(numerator, denominator)
	numerator /= gcd
	denominator /= gcd
	return D128{
		sign:        sign,
		denominator: denominator,
		numerator:   numerator,
	}
}

func (a D128) Div(b D128) D128 {
	b.denominator, b.numerator = b.numerator, b.denominator
	return a.Mul(b)
}

func (a D128) Cmp(b D128) Compare {
	if a.sign != b.sign {
		if a.sign == Positive {
			return Greater
		}
		return Less
	}
	gcd := Gcd64(a.denominator, b.denominator)
	a.denominator /= gcd
	b.denominator /= gcd
	a.numerator *= b.denominator
	b.numerator *= a.denominator
	if a.numerator > b.numerator {
		return Greater
	} else if a.numerator < b.numerator {
		return Less
	}
	return Equal
}

func (a D128) Float32() float32 {
	f := float32(a.numerator) / float32(a.denominator)
	if a.sign == Negative {
		f = -f
	}
	return f
}

func (a D128) Float64() float64 {
	f := float64(a.numerator) / float64(a.denominator)
	if a.sign == Negative {
		f = -f
	}
	return f
}
