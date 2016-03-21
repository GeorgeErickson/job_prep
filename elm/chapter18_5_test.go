package elm

import (
	"testing"
	"testing/quick"

	"github.com/stretchr/testify/require"
)

func BruteThreeSum(a []int, x int) bool {
	// O(n^3) time, O(1) space.
	for _, v1 := range a {
		for _, v2 := range a {
			for _, v3 := range a {
				if v1+v2+v3 == x {
					return true
				}
			}
		}
	}

	return false
}

func ThreeSum(a []int, x int) bool {
	// O(n^2) time, O(n) space.
	hsh := make(map[int]struct{})
	for _, v1 := range a {

		for _, v2 := range a {
			v := x - (v1 + v2)
			hsh[v] = struct{}{}
		}
	}

	for _, v := range a {
		if _, found := hsh[v]; found {
			return true
		}
	}

	return false
}

func TestThreeSum(t *testing.T) {
	assert := require.New(t)
	example := []int{11, 2, 5, 7, 3}

	assert.True(ThreeSum(example, 21))

	if err := quick.CheckEqual(BruteThreeSum, ThreeSum, nil); err != nil {
		t.Error(err)
	}
}
