package ds

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMatrix(t *testing.T) {
	assert := require.New(t)

	m := MatrixFrom([][]int{
		{1},
		{2},
	})

	assert.Equal(m.nCol, 1)
	assert.Equal(m.nRow, 2)

	out := m.Product(MatrixFrom([][]int{{1, 5}}))

	assert.Equal(out.nCol, 2)
	assert.Equal(out.nRow, 2)

	assert.True(out.Equals([][]int{
		{1, 5},
		{2, 10},
	}))
}

func TestMatrixSquare(t *testing.T) {
	assert := require.New(t)

	m := MatrixFrom([][]int{
		{0, 1},
		{1, 1},
	})
	sq := m.Square()
	assert.True(sq.Equals([][]int{
		{1, 1},
		{1, 2},
	}))

	assert.True(sq.Square().Equals([][]int{
		{2, 3},
		{3, 5},
	}))

}
