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

		if b.dom(pIdx, vIdx) {
			return
		}
		b.swap(pIdx, vIdx)
		vIdx = pIdx
	}

}

func (b *Binary) dom(p, j int) bool {
	// no child exists.
	if j >= len(b.data) {
		return true
	}

	return b.dominates(b.data[p], b.data[j])
}

func (b *Binary) swap(i, j int) {
	b.data[i], b.data[j] = b.data[j], b.data[i]
}

// Remove/Return root and reorg
func (b *Binary) Extract() int {
	if len(b.data) == 0 {
		return 0
	}
	dominateVal := b.data[0]
	b.data[0], b.data = b.data[len(b.data)-1], b.data[:len(b.data)-1]
	pIdx := 0

	for {
		lIdx := pIdx * 2
		rIdx := lIdx + 1

		if b.dom(pIdx, lIdx) && b.dom(pIdx, rIdx) {
			break
		}

		lExists := lIdx < len(b.data)
		rExists := rIdx < len(b.data)

		if lExists && rExists {
			l := b.data[lIdx]
			r := b.data[rIdx]

			if b.dom(l, r) {
				b.swap(lIdx, pIdx)
				pIdx = lIdx
			} else {
				b.swap(rIdx, pIdx)
				pIdx = rIdx
			}
		} else if lExists {
			b.swap(lIdx, pIdx)
			pIdx = lIdx
		} else if rExists {
			b.swap(rIdx, pIdx)
			pIdx = rIdx
		}

	}

	return dominateVal
}
