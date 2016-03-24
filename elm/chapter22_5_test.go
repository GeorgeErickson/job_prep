package elm

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func LongestIncSubArray(a []int) (int, int) {

	cl, maxL, maxLIdx := 1, 1, 0

	for i, v := range a {
		if i == 0 {
			continue
		}

		if v > a[i-1] {
			cl++
			if cl > maxL {
				maxL = cl
				maxLIdx = i
			}
		} else {
			cl = 1
		}
	}

	return maxLIdx - maxL + 1, maxLIdx
}

func TestLongestIncSubArray(t *testing.T) {
	cases := []struct {
		a    []int
		s, e int
	}{
		{[]int{2, 11, 3, 5, 13, 7, 19, 17, 23}, 2, 4},
		{[]int{}, 0, 0},
	}
	assert := require.New(t)
	for _, tc := range cases {
		s, e := LongestIncSubArray(tc.a)

		assert.Equal(tc.s, s)
		assert.Equal(tc.e, e)
	}

}
