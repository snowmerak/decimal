package decimal

import "fmt"

type D64 struct {
	sign        Sign
	numerator   uint32
	denominator uint32
}

func New64(sign Sign, numerator, denominator uint32) D64 {
	gcd := Gcd32(numerator, denominator)
	numerator /= gcd
	denominator /= gcd
	return D64{
		sign:        sign,
		numerator:   numerator,
		denominator: denominator,
	}
}

func (a D64) String() string {
	sign := ""
	if a.sign == Negative {
		sign = "-"
	}
	return fmt.Sprintf("%s%d/%d", sign, a.numerator, a.denominator)
}

func (a D64) Add(b D64) D64 {
	if a.sign != b.sign {
		if a.sign == Negative {
			a.sign = Positive
			return b.Sub(a)
		}
		b.sign = Positive
		return a.Sub(b)
	}
	gcd := Gcd32(a.denominator, b.denominator)
	a.denominator /= gcd
	b.denominator /= gcd
	a.numerator *= b.denominator
	b.numerator *= a.denominator
	denominator := a.denominator * b.denominator * gcd
	numerator := a.numerator + b.numerator
	gcd = Gcd32(numerator, denominator)
	numerator /= gcd
	denominator /= gcd
	return D64{
		sign:        a.sign,
		denominator: denominator,
		numerator:   numerator,
	}
}

func (a D64) Sub(b D64) D64 {
	if a.sign != b.sign {
		if a.sign == Positive {
			a.sign = Negative
			return a.Add(b)
		}
		b.sign = Negative
		return b.Add(a)
	}
	gcd := Gcd32(a.denominator, b.denominator)
	a.denominator /= gcd
	b.denominator /= gcd
	a.numerator *= b.denominator
	b.numerator *= a.denominator
	denominator := a.denominator * b.denominator * gcd
	numerator := uint32(0)
	sign := Positive
	if a.numerator >= b.numerator {
		numerator = a.numerator - b.numerator
	} else {
		sign = Negative
		numerator = b.numerator - a.numerator
	}
	if numerator == 0 {
		denominator = 1
	}
	gcd = Gcd32(numerator, denominator)
	numerator /= gcd
	denominator /= gcd
	return D64{
		sign:        sign,
		denominator: denominator,
		numerator:   numerator,
	}
}

func (a D64) Mul(b D64) D64 {
	sign := Positive
	if a.sign != b.sign {
		sign = Negative
	}
	gcd1 := Gcd32(a.numerator, b.denominator)
	gcd2 := Gcd32(b.numerator, a.denominator)
	numerator := (a.numerator / gcd1) * (b.numerator / gcd2)
	denominator := (a.denominator / gcd2) * (b.denominator / gcd1)
	gcd := Gcd32(numerator, denominator)
	numerator /= gcd
	denominator /= gcd
	return D64{
		sign:        sign,
		denominator: denominator,
		numerator:   numerator,
	}
}

func (a D64) Div(b D64) D64 {
	b.denominator, b.numerator = b.numerator, b.denominator
	return a.Mul(b)
}

func (a D64) Cmp(b D64) Compare {
	if a.sign != b.sign {
		if a.sign == Positive {
			return Greater
		}
		return Less
	}
	gcd := Gcd32(a.denominator, b.denominator)
	a.denominator /= gcd
	b.denominator /= gcd
	a.numerator *= b.denominator
	b.numerator *= a.denominator
	if a.numerator > b.numerator {
		return Greater
	}
	if a.numerator < b.numerator {
		return Less
	}
	return Equal
}
