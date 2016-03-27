package btree

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func AssertSorted(t *testing.T, a []Value) {
	if len(a) == 0 {
		return
	}

	p := a[0]
	for i := 1; i < len(a); i++ {
		if a[i] < p {
			t.Fatalf("Array %v should be sorted.", a)
			return
		}

		p = a[i]
	}
}

func exampleTree() *Node {
	return &Node{19,
		&Node{7,
			&Node{3,
				&Node{2, nil, nil},
				&Node{5, nil, nil},
			},
			&Node{11,
				nil,
				&Node{17,
					&Node{13, nil, nil},
					nil,
				},
			},
		},
		&Node{43,
			&Node{23,
				nil,
				&Node{37,
					&Node{29, nil, &Node{31, nil, nil}},
					&Node{41, nil, nil},
				},
			},
			&Node{47,
				nil,
				&Node{53, nil, nil},
			},
		},
	}
}

func TestBtree(t *testing.T) {
	// page 250
	bt := exampleTree()

	AssertSorted(t, bt.Walk(InOrder))
	assert := require.New(t)
	assert.Equal(5, bt.Height())
}

func TestBtreeHeight(t *testing.T) {
	// page 250
	bt := &Node{47,
		nil,
		&Node{53, nil, nil},
	}
	assert := require.New(t)
	assert.Equal(1, bt.Height())

	assert.Equal(0, bt.r.Height())
}

func TestBtreeIsBalanced(t *testing.T) {
	bt := &Node{47,
		nil,
		&Node{53, nil, nil},
	}
	assert := require.New(t)
	assert.Equal(1, bt.Height())
	assert.True(bt.IsBalanced())
	bt.r.r = &Node{65, nil, nil}
	assert.False(bt.IsBalanced())
}
