package elm

import (
	"math"
	"testing"

	"github.com/stretchr/testify/require"
)

// x^y don't worry about overflow
func Pow(x, y int) int {
	switch {
	case y == 0 || x == 1:
		return 1
	case x == 0:
		return 0
	}

	r := 1
	for y > 1 {
		if y&1 == 1 {
			y--
			r *= x
		}
		x *= x
		y >>= 1
	}

	return r * x
}

func mathPow(x, y int) int {
	return int(math.Pow(float64(x), float64(y)))
}

func TestExp(t *testing.T) {
	assert := require.New(t)
	for y := 0; y < 20; y++ {
		assert.Equal(mathPow(3, y), Pow(3, y))
	}

}
