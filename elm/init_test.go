package elm

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNChooseK(t *testing.T) {
	assert := require.New(t)
	assert.Equal(4950, NChooseK(100, 2))
	assert.Equal(100, NChooseK(100, 99))
	assert.Equal(1, NChooseK(100, 100))
}
