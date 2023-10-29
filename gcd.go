package decimal

import "math/bits"

func SlowGcd[T ~uint8 | ~uint16 | ~uint32 | ~uint64](u, v T) T {
	for v != 0 {
		u, v = v, u%v
	}
	return u
}

func Gcd64(u, v uint64) uint64 {
	var shift int
	if u == 0 {
		return v
	}
	if v == 0 {
		return u
	}
	shift = bits.TrailingZeros64(uint64(u | v))
	u >>= bits.TrailingZeros64(uint64(u))
	for {
		var m uint64
		v >>= bits.TrailingZeros64(uint64(v))
		v -= u
		m = uint64(int64(v) >> 63)
		u += v & m
		v = (v + m) ^ m
		if v == 0 {
			break
		}
	}
	return u << shift
}

func Gcd32(u, v uint32) uint32 {
	var shift int
	if u == 0 {
		return v
	}
	if v == 0 {
		return u
	}
	shift = bits.TrailingZeros32(uint32(u | v))
	u >>= bits.TrailingZeros32(uint32(u))
	for {
		var m uint32
		v >>= bits.TrailingZeros32(uint32(v))
		v -= u
		m = uint32(int32(v) >> 31)
		u += v & m
		v = (v + m) ^ m
		if v == 0 {
			break
		}
	}
	return u << shift
}

func Gcd16(u, v uint16) uint16 {
	var shift int
	if u == 0 {
		return v
	}
	if v == 0 {
		return u
	}
	shift = bits.TrailingZeros16(uint16(u | v))
	u >>= bits.TrailingZeros16(uint16(u))
	for {
		var m uint16
		v >>= bits.TrailingZeros16(uint16(v))
		v -= u
		m = uint16(int16(v) >> 15)
		u += v & m
		v = (v + m) ^ m
		if v == 0 {
			break
		}
	}
	return u << shift
}

func Gcd8(u, v uint8) uint8 {
	var shift int
	if u == 0 {
		return v
	}
	if v == 0 {
		return u
	}
	shift = bits.TrailingZeros8(uint8(u | v))
	u >>= bits.TrailingZeros8(uint8(u))
	for {
		var m uint8
		v >>= bits.TrailingZeros8(uint8(v))
		v -= u
		m = uint8(int8(v) >> 7)
		u += v & m
		v = (v + m) ^ m
		if v == 0 {
			break
		}
	}
	return u << shift
}
