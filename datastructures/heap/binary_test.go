package heap

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestBinary(t *testing.T) {
	b := &MinBinary{}
	assert := require.New(t)
	b.Insert(7)
	assert.Equal([]int64{7}, b.data)
	b.Insert(5)
	assert.Equal([]int64{5, 7}, b.data)
	b.Insert(4)
	assert.Equal([]int64{4, 5, 7}, b.data)
	b.Insert(6)
	assert.Equal([]int64{4, 5, 7, 6}, b.data)

	assert.Equal(b.Pop(), int64(4))
	assert.Equal([]int64{5, 6, 7}, b.data)
}
