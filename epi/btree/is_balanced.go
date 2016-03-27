package btree

// IsBalancedBrute returns true, if for each node in the tree, the diff in height between its left and right children is <= 1
func (n *Node) IsBalancedBrute() bool {
	b := true

	n.Each(func(c *Node) {
		dt := c.l.childHeight() - c.r.childHeight()

		if dt > 1 || dt < -1 {
			b = false
		}
	})
	return b
}

func (n *Node) isB() (int, bool) {
	if n == nil {
		return -1, true
	}
	lh, lb := n.l.isB()
	if !lb {
		return 0, false
	}

	rh, rb := n.r.isB()
	if !rb {
		return 0, false
	}

	return max(lh, rh) + 1, abs(lh-rh) <= 1
}

func (n *Node) IsBalanced() bool {
	_, b := n.isB()
	return b
}
