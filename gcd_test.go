package decimal_test

import (
	"testing"

	"github.com/snowmerak/decimal"
)

func FuzzGcd64(f *testing.F) {
	f.Add(uint64(23), uint64(37))
	f.Fuzz(func(t *testing.T, u uint64, v uint64) {
		fast := decimal.Gcd64(u, v)
		slow := decimal.SlowGcd(u, v)
		if fast != slow {
			t.Logf("u: %d, v: %d, fast(%d) != slow(%d)", u, v, fast, slow)
			t.Fail()
		}
	})
}

func FuzzGcd32(f *testing.F) {
	f.Add(uint32(23), uint32(37))
	f.Fuzz(func(t *testing.T, u uint32, v uint32) {
		fast := decimal.Gcd32(u, v)
		slow := decimal.SlowGcd(u, v)
		if fast != slow {
			t.Logf("u: %d, v: %d, fast(%d) != slow(%d)", u, v, fast, slow)
			t.Fail()
		}
	})
}

func FuzzGcd16(f *testing.F) {
	f.Add(uint16(23), uint16(37))
	f.Fuzz(func(t *testing.T, u uint16, v uint16) {
		fast := decimal.Gcd16(u, v)
		slow := decimal.SlowGcd(u, v)
		if fast != slow {
			t.Logf("u: %d, v: %d, fast(%d) != slow(%d)", u, v, fast, slow)
			t.Fail()
		}
	})
}

func FuzzGcd8(f *testing.F) {
	f.Add(uint8(23), uint8(37))
	f.Fuzz(func(t *testing.T, u uint8, v uint8) {
		fast := decimal.Gcd8(u, v)
		slow := decimal.SlowGcd(u, v)
		if fast != slow {
			t.Logf("u: %d, v: %d, fast(%d) != slow(%d)", u, v, fast, slow)
			t.Fail()
		}
	})
}
