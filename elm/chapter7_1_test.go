package elm

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func ReverseRune(inp []rune) string {
	n := len(inp)
	for i := 0; i < n/2; i++ {
		j := n - i - 1
		inp[i], inp[j] = inp[j], inp[i]
	}
	return string(inp)
}

func IntToString(v int) string {
	isNeg := false
	if v < 0 {
		isNeg = true
		v *= -1
	}
	if v == 0 {
		return "0"
	}

	var out []rune

	for v > 0 {
		out = append(out, '0'+rune(v%10))
		v /= 10
	}

	if isNeg {
		out = append(out, '-')
	}

	return ReverseRune(out)
}

func StringToInt(s string) int {
	m := 1
	b := 10
	runes := []rune(s)
	neg := 1
	if runes[0] == '-' {
		runes = runes[1:]
		neg = -1
	}

	out := 0
	for i := len(runes) - 1; i >= 0; i-- {
		v := runes[i]
		out += int(v-'0') * m
		m *= b
	}

	return out * neg
}

func TestIntToString(t *testing.T) {
	assert := require.New(t)
	for _, v := range []int{-100, 231, 0, 1, 300, 1234} {
		act := fmt.Sprintf("%v", v)
		assert.Equal(act, IntToString(v))
		assert.Equal(v, StringToInt(act))
	}

}
