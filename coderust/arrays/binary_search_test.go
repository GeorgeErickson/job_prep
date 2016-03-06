package arrays

import "testing"

func TestBinarySearch(t *testing.T) {
	a := []int{1, 10, 20, 47, 59, 63, 75, 88, 99}

	for i, k := range a {
		act := BinarySearch(a, k)

		if i != act {
			t.Errorf("BinarySearch(%v, %v) should equal %v, not %v", a, k, i, act)
		}
	}
}
