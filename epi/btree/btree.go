package btree

type Value int

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

type Node struct {
	V Value
	l *Node
	r *Node
}

type WalkFn func(n *Node, out *[]Value)

func (n *Node) Height() int {
	if n == nil {
		return -1
	}

	return 1 + max(n.l.Height(), n.r.Height())
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
