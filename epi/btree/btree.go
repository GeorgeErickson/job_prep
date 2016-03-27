package btree

type Value int

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func abs(a int) int {
	if a < 0 {
		a *= -1
	}
	return a
}

type Node struct {
	V Value
	l *Node
	r *Node
}

func (n *Node) IsLeaf() bool {
	return n.l == nil && n.r == nil
}

type WalkFn func(n *Node, out *[]Value)

func (n *Node) childHeight() int {
	if n == nil {
		return 0
	}

	return n.Height() + 1
}

func (n *Node) Height() int {

	return max(n.l.childHeight(), n.r.childHeight())
}

func (n *Node) Each(fn func(n *Node)) {
	if n == nil {
		return
	}

	n.l.Each(fn)
	fn(n)
	n.r.Each(fn)
}

func InOrder(n *Node, out *[]Value) {
	if n == nil {
		return
	}

	InOrder(n.l, out)
	*out = append(*out, n.V)
	InOrder(n.r, out)
}

func PreOrder(n *Node, out *[]Value) {
	if n == nil {
		return
	}
	*out = append(*out, n.V)
	InOrder(n.l, out)
	InOrder(n.r, out)
}

func (n *Node) Walk(fn WalkFn) []Value {
	var out []Value
	fn(n, &out)
	return out
}
