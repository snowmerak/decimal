package decimal_test

import (
	"testing"

	"github.com/snowmerak/decimal"
)

func TestD128Add(t *testing.T) {
	cases := []Case[decimal.D128, string]{
		{Input1: decimal.New128(decimal.Positive, 1, 2), Input2: decimal.New128(decimal.Positive, 1, 2), Expect: "1/1"},
		{Input1: decimal.New128(decimal.Positive, 1, 2), Input2: decimal.New128(decimal.Positive, 1, 3), Expect: "5/6"},
		{Input1: decimal.New128(decimal.Positive, 1, 2), Input2: decimal.New128(decimal.Positive, 1, 4), Expect: "3/4"},
		{Input1: decimal.New128(decimal.Positive, 1, 2), Input2: decimal.New128(decimal.Positive, 1, 5), Expect: "7/10"},
		{Input1: decimal.New128(decimal.Positive, 1, 2), Input2: decimal.New128(decimal.Positive, 1, 6), Expect: "2/3"},
		{Input1: decimal.New128(decimal.Positive, 5, 6), Input2: decimal.New128(decimal.Positive, 3, 7), Expect: "53/42"},
		{Input1: decimal.New128(decimal.Negative, 1, 2), Input2: decimal.New128(decimal.Positive, 1, 2), Expect: "0/1"},
		{Input1: decimal.New128(decimal.Negative, 1, 6), Input2: decimal.New128(decimal.Positive, 1, 3), Expect: "1/6"},
		{Input1: decimal.New128(decimal.Negative, 3, 65), Input2: decimal.New128(decimal.Positive, 7, 16), Expect: "407/1040"},
		{Input1: decimal.New128(decimal.Negative, 9, 991), Input2: decimal.New128(decimal.Negative, 1033, 5), Expect: "-1023748/4955"},
		{Input1: decimal.New128(decimal.Negative, 1035, 1), Input2: decimal.New128(decimal.Negative, 666, 3), Expect: "-1257/1"},
	}

	for _, c := range cases {
		if c.Input1.Add(c.Input2).String() != c.Expect {
			t.Logf("Input1: %s, Input2: %s, Expect: %s, Actual: %s", c.Input1, c.Input2, c.Expect, c.Input1.Add(c.Input2))
			t.Fail()
		}
	}
}

func TestD128Mul(t *testing.T) {
	cases := []Case[decimal.D128, string]{
		{Input1: decimal.New128(decimal.Positive, 1, 2), Input2: decimal.New128(decimal.Positive, 1, 2), Expect: "1/4"},
		{Input1: decimal.New128(decimal.Positive, 1, 2), Input2: decimal.New128(decimal.Positive, 1, 3), Expect: "1/6"},
		{Input1: decimal.New128(decimal.Positive, 1, 2), Input2: decimal.New128(decimal.Positive, 1, 4), Expect: "1/8"},
		{Input1: decimal.New128(decimal.Positive, 1, 2), Input2: decimal.New128(decimal.Positive, 1, 5), Expect: "1/10"},
		{Input1: decimal.New128(decimal.Positive, 1, 2), Input2: decimal.New128(decimal.Positive, 1, 6), Expect: "1/12"},
		{Input1: decimal.New128(decimal.Negative, 5, 12), Input2: decimal.New128(decimal.Positive, 3, 7), Expect: "-5/28"},
		{Input1: decimal.New128(decimal.Negative, 7, 2), Input2: decimal.New128(decimal.Negative, 8, 3), Expect: "28/3"},
	}

	for _, c := range cases {
		if c.Input1.Mul(c.Input2).String() != c.Expect {
			t.Logf("Input1: %s, Input2: %s, Expect: %s, Actual: %s", c.Input1, c.Input2, c.Expect, c.Input1.Mul(c.Input2))
			t.Fail()
		}
	}
}

func TestD128Div(t *testing.T) {
	cases := []Case[decimal.D128, string]{
		{Input1: decimal.New128(decimal.Positive, 1, 2), Input2: decimal.New128(decimal.Positive, 1, 2), Expect: "1/1"},
		{Input1: decimal.New128(decimal.Positive, 1, 2), Input2: decimal.New128(decimal.Positive, 1, 3), Expect: "3/2"},
		{Input1: decimal.New128(decimal.Positive, 1, 2), Input2: decimal.New128(decimal.Positive, 1, 4), Expect: "2/1"},
		{Input1: decimal.New128(decimal.Positive, 1, 2), Input2: decimal.New128(decimal.Positive, 1, 5), Expect: "5/2"},
		{Input1: decimal.New128(decimal.Positive, 1, 2), Input2: decimal.New128(decimal.Positive, 1, 6), Expect: "3/1"},
		{Input1: decimal.New128(decimal.Negative, 5, 12), Input2: decimal.New128(decimal.Positive, 3, 7), Expect: "-35/36"},
		{Input1: decimal.New128(decimal.Negative, 7, 2), Input2: decimal.New128(decimal.Negative, 8, 3), Expect: "21/16"},
	}

	for _, c := range cases {
		if c.Input1.Div(c.Input2).String() != c.Expect {
			t.Logf("Input1: %s, Input2: %s, Expect: %s, Actual: %s", c.Input1, c.Input2, c.Expect, c.Input1.Div(c.Input2))
			t.Fail()
		}
	}
}

func TestD128Comp(t *testing.T) {
	cases := []Case[decimal.D128, decimal.Compare]{
		{Input1: decimal.New128(decimal.Positive, 1, 2), Input2: decimal.New128(decimal.Positive, 1, 2), Expect: decimal.Equal},
		{Input1: decimal.New128(decimal.Positive, 1, 2), Input2: decimal.New128(decimal.Positive, 1, 3), Expect: decimal.Greater},
		{Input1: decimal.New128(decimal.Positive, 1, 2), Input2: decimal.New128(decimal.Positive, 1, 4), Expect: decimal.Greater},
		{Input1: decimal.New128(decimal.Positive, 1, 2), Input2: decimal.New128(decimal.Positive, 1, 5), Expect: decimal.Greater},
		{Input1: decimal.New128(decimal.Positive, 1, 2), Input2: decimal.New128(decimal.Positive, 1, 6), Expect: decimal.Greater},
		{Input1: decimal.New128(decimal.Negative, 5, 12), Input2: decimal.New128(decimal.Positive, 3, 7), Expect: decimal.Less},
		{Input1: decimal.New128(decimal.Negative, 7, 2), Input2: decimal.New128(decimal.Negative, 8, 3), Expect: decimal.Greater},
	}

	for _, c := range cases {
		cmp := c.Input1.Cmp(c.Input2)
		if cmp != c.Expect {
			t.Logf("Input1: %s, Input2: %s, Expect: %v, Actual: %v", c.Input1, c.Input2, c.Expect, cmp)
			t.Fail()
		}
	}
}
