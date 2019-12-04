package day4

import (
	"testing"
)

type test struct {
	in  int
	out bool
}

func TestStandard(t *testing.T)  {
	tests := []test{
		{in: 679999, out: true,},
		{in: 144559, out: true,},
		{in: 144689, out: true,},
		{in: 145688, out: true,},
		{in: 111111, out: true,},
		{in: 144888, out: true,},
		{in: 145888, out: true,},
		{in: 145882, out: false,},
	}

	for _, tc := range tests {
		result := isMatchingStandardRules(tc.in)
		if result != tc.out {
			t.Fatalf("expected: %v, got: %v for %v", tc.out, result, tc.in)
		}
	}
}

func TestBonus(t *testing.T) {
	type test struct {
		in  int
		out bool
	}

	tests := []test{
		{in: 679999, out: false,},
		{in: 144559, out: true,},
		{in: 144689, out: true,},
		{in: 145688, out: true,},
		{in: 111111, out: false,},
		{in: 144888, out: true,},
		{in: 145888, out: false,},
		{in: 145881, out: false,},
	}

	for _, tc := range tests {
		result := isMatchingBonusRules(tc.in)
		if result != tc.out {
			t.Fatalf("expected: %v, got: %v for %v", tc.out, result, tc.in)
		}
	}
}
