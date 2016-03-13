package elm

import "testing"

func Parity(v int) int {
	i := 0
	for v > 0 {
		i ^= 1
		v &= (v - 1) // clear lowest set bit
	}

	return i
}

func TestParity(t *testing.T) {
	zeros := []int{0, 3, 5, 6}
	ones := []int{1, 2, 4, 7, 8}

	for _, v := range zeros {
		if Parity(v) != 0 {
			t.Errorf("%s should have a parity of 0", v)
		}
	}

	for _, v := range ones {
		if Parity(v) != 1 {
			t.Errorf("%s should have a parity of 0", v)
		}
	}

}
