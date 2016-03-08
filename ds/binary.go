package ds

type MinBinary struct {
	data []int
}

func NewMinBinary(items ...int) *MinBinary {
	b := &MinBinary{data: items}

	for i := b.lastIdx(); i >= 0; i-- {
		b.bubbleDown(i)
	}

	return b
}

func (b *MinBinary) Sorted() []int {
	n := len(b.data)

	out := make([]int, n)
	for i := 0; i < n; i++ {
		out[i] = b.Pop()
	}

	return out
}

func (b *MinBinary) parentIdx(i int) int {
	if i == 0 {
		return -1
	}

	return i / 2
}

func (b *MinBinary) lastIdx() int {
	return len(b.data) - 1
}

func (b *MinBinary) childIdx(i int) int {
	return (i * 2) + 1
}

func (b *MinBinary) swap(i, j int) {
	b.data[i], b.data[j] = b.data[j], b.data[i]
}

func (b *MinBinary) Pop() int {
	if len(b.data) == 0 {
		return -1
	}
	min := b.data[0]
	b.swap(0, b.lastIdx())
	b.data = b.data[:b.lastIdx()]
	b.bubbleDown(0)
	return min
}

func (b *MinBinary) bubbleDown(pi int) {
	ci := b.childIdx(pi)
	minIdx := pi

	// swap with most dominate child.
	for i := 0; i <= 1; i++ {
		if ci+i > b.lastIdx() {
			break
		}

		if b.data[minIdx] > b.data[ci+i] {
			minIdx = ci + i
		}
	}

	if pi != minIdx {
		b.swap(pi, minIdx)
		b.bubbleDown(minIdx)
	}
}

func (b *MinBinary) Insert(x int) {
	b.data = append(b.data, x)
	b.bubbleUp(b.lastIdx())
}

func (b *MinBinary) bubbleUp(i int) {
	pi := b.parentIdx(i)

	// at root
	if pi == -1 {
		return
	}

	// child dominates parent
	if b.data[pi] > b.data[i] {
		b.swap(pi, i)
		b.bubbleUp(pi)
	}
}
