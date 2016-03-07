package heap

// DominateFunc returns true if x dominates y
type DominateFunc func(x, y int) bool

func Min(x, y int) bool {
	return x <= y
}

type Binary struct {
	dominates DominateFunc
	data      []int
}

func NewMinBinary() *Binary {
	return &Binary{
		dominates: Min,
		data:      []int{},
	}
}

func (b *Binary) Insert(v int) {
	b.data = append(b.data, v)
	vIdx := len(b.data) - 1
	pIdx := vIdx

	for pIdx > 0 {
		pIdx = pIdx / 2
		p := b.data[pIdx]

		if b.dominates(p, v) {
			return
		}

		b.data[vIdx], b.data[pIdx] = b.data[pIdx], b.data[vIdx]
		vIdx = pIdx
	}

}
