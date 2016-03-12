package cracking

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSquare33(t *testing.T) {
	assert := require.New(t)
	m := [3][3]int{
		{0, 1, 0},
		{0, 0, 1},
		{1, 1, 1},
	}
	out := square33(m)
	exp := [3][3]int{
		{0, 0, 1},
		{1, 1, 1},
		{1, 2, 2},
	}
	assert.Equal(exp, out)

	for i, exp := range []int{0, 1, 1, 2, 3, 5, 8, 13, 21, 34} {
		act := fib2(i)
		assert.Equal(exp, act, "fib(%v) should equal(%v) got %v", i, exp, act)
		// fmt.Println(i, fib2(i), exp)
	}
	// assert.Equal(fib2(10), 55)
	// fib(11)
}

func BenchmarkFib(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fib(100000)
	}
}

func BenchmarkFib2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fib2(100000)
	}
}

func TestPowerBySquaring(t *testing.T) {
	fmt.Println(fib2(100000))
}
