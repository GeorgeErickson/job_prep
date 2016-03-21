package elm

import (
	"testing"
	"testing/quick"

	"github.com/stretchr/testify/require"
)

type IntList []int

func BruteMaxSubarrayValue(a []int16) int {
	max := 0
	// iterate through each subarray
	for w := 1; w <= len(a); w++ {
		numW := len(a) - w + 1
		for o := 0; o < numW; o++ {
			lm := 0
			for _, v := range a[o : o+w] {
				lm += int(v)
			}

			if lm > max {
				max = lm
			}
		}
	}
	return max
}

func MaxSubarrayValue(a []int16) int {
	max := 0

	me := 0
	for _, val := range a {
		v := int(val)
		me += v
		if v > me {
			me = v
		}

		if me > max {
			max = me
		}
	}

	return max
}

func TestMaxSubArray(t *testing.T) {
	assert := require.New(t)

	assert.Equal(1479, MaxSubarrayValue([]int16{904, 40, 523, 12, -335, -385, -124, 481, -31}))

	if err := quick.CheckEqual(MaxSubarrayValue, BruteMaxSubarrayValue, nil); err != nil {
		t.Error(err)
	}
}
