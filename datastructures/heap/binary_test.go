package heap

import (
	"math/rand"
	"sort"
	"testing"

	"github.com/stretchr/testify/require"
)

func randomIntList(n int) []int {
	out := make([]int, n)
	for i := 0; i < n; i++ {
		out[i] = rand.Int()
	}

	return out
}

func TestBinary(t *testing.T) {
	b := NewMinBinary()
	assert := require.New(t)
	b.Insert(7)
	assert.Equal([]int{7}, b.data)
	b.Insert(5)
	assert.Equal([]int{5, 7}, b.data)
	b.Insert(4)
	assert.Equal([]int{4, 5, 7}, b.data)
	b.Insert(6)
	assert.Equal([]int{4, 5, 7, 6}, b.data)

	assert.Equal(b.Pop(), 4)
	assert.Equal([]int{5, 6, 7}, b.data)

	b = NewMinBinary(randomIntList(100)...)
	assert.True(sort.IntsAreSorted(b.Sorted()))
}
