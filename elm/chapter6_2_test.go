package elm

import (
	"reflect"
	"testing"
)

// Increment
func increment62(v []uint) []uint {
	for i := len(v) - 1; i >= 0; i-- {
		if v[i] < 9 {
			v[i]++

			return v

		}
		v[i] = 0

	}

	return append([]uint{1}, v...)
}

func addDigits(v1 []int, v2 []int) []int {
	top, bot := v1, v2

	if len(bot) > len(top) {
		top, bot = bot, top
	}

	out := make([]int, len(top)+1)

	for i := 1; i <= len(top); i++ {
		ti := len(top) - i
		bi := len(bot) - i
		oi := len(out) - i

		t := top[ti]
		b := 0
		if bi >= 0 {
			b = bot[bi]
		}

		out[oi] += t + b

		if out[oi] > 9 {
			out[oi-1] += out[oi] / 10
			out[oi] = out[oi] % 10
		}
	}

	if out[0] == 0 {
		return out[1:]
	}

	return out
}

func multiplyDigits(v1 []int, v2 []int) []int {
	top, bot := v1, v2

	if len(bot) > len(top) {
		top, bot = bot, top
	}

	out := make([]int, len(top)+1)

	for b, shift := len(bot)-1, 0; b >= 0; b-- {
		bv := bot[b]

		part := make([]int, len(top))
		for t := len(top) - 1; t >= 0; t-- {
			part[t] = part[t] + top[t]*bv

			if part[t] > 9 {
				ni := t - 1
				nv := part[t] / 10
				part[t] %= 10
				if ni >= 0 {
					part[ni] = nv
				} else {
					part = append([]int{nv}, part...)
				}

			}
		}
		for i := 0; i < shift; i++ {
			part = append(part, 0)
		}

		out = addDigits(out, part)

		shift++
	}

	if out[0] == 0 {
		return out[1:]
	}

	return out
}

type addDigitsTestCase struct {
	v1       []int
	v2       []int
	expected []int
}

func TestAddDigits(t *testing.T) {
	cases := []addDigitsTestCase{
		{[]int{0}, []int{1}, []int{1}},
		{[]int{9, 9}, []int{1}, []int{1, 0, 0}},
	}

	for _, tc := range cases {
		actual := addDigits(tc.v1, tc.v2)
		if !reflect.DeepEqual(actual, tc.expected) {
			t.Errorf("addDigits(%v, %v) should return %v, not %v", tc.v1, tc.v2, tc.expected, actual)
		}

		actual2 := addDigits(tc.v2, tc.v1)
		if !reflect.DeepEqual(actual2, tc.expected) {
			t.Errorf("addDigits(%v, %v) should return %v, not %v", tc.v2, tc.v1, tc.expected, actual)
		}
	}
}

func TestMulDigits(t *testing.T) {
	cases := []addDigitsTestCase{
		{[]int{0}, []int{1}, []int{0}},
		{[]int{9, 9}, []int{2}, []int{1, 9, 8}},
		{[]int{9, 9}, []int{9, 9, 9}, []int{9, 8, 9, 0, 1}},
	}

	for _, tc := range cases {
		actual := multiplyDigits(tc.v1, tc.v2)
		if !reflect.DeepEqual(actual, tc.expected) {
			t.Errorf("addDigits(%v, %v) should return %v, not %v", tc.v1, tc.v2, tc.expected, actual)
		}

		actual2 := multiplyDigits(tc.v2, tc.v1)
		if !reflect.DeepEqual(actual2, tc.expected) {
			t.Errorf("addDigits(%v, %v) should return %v, not %v", tc.v2, tc.v1, tc.expected, actual)
		}
	}
}

type incrementTestCase struct {
	inp      []uint
	expected []uint
}

func TestIncrement62(t *testing.T) {
	cases := []incrementTestCase{
		{[]uint{0}, []uint{1}},
		{[]uint{9}, []uint{1, 0}},
		{[]uint{1, 2, 9}, []uint{1, 3, 0}},
		{[]uint{9, 9, 9}, []uint{1, 0, 0, 0}},
	}

	for _, tc := range cases {
		actual := increment62(tc.inp)
		if !reflect.DeepEqual(actual, tc.expected) {
			t.Errorf("increment(%v) should return %v, not %v", tc.inp, tc.expected, actual)
		}
	}
}
