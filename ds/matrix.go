package ds

import "fmt"

type Matrix struct {
	data []int
	nRow int
	nCol int
}

func MatrixFrom(data [][]int) *Matrix {
	m := NewMatrix(len(data), len(data[0]))
	m.Reset(data)
	return m
}

func (m *Matrix) Row(i int) []int {
	rs := i * m.nCol
	return m.data[rs : rs+m.nCol]
}

func (m *Matrix) Print() {
	for i := 0; i < m.nRow; i++ {
		fmt.Printf("%v\n", m.Row(i))
	}
}

func NewMatrix(nRow, nCol int) *Matrix {
	return &Matrix{
		nRow: nRow,
		nCol: nCol,
		data: make([]int, nRow*nCol),
	}

}

func (m *Matrix) Reset(data [][]int) {
	for r, cols := range data {
		for c, v := range cols {
			m.Set(r, c, v)
		}
	}
}

func (m *Matrix) Set(r, c, v int) {
	m.data[(m.nCol*r)+c] = v
}

func (m *Matrix) Get(r, c int) int {
	return m.data[(m.nCol*r)+c]
}

func (m *Matrix) Square() *Matrix {
	return m.Product(m)
}

func (m *Matrix) Product(b *Matrix) *Matrix {
	out := NewMatrix(m.nRow, b.nCol)
	if m.nCol != b.nRow {
		panic("can't multiply")
	}

	for r := 0; r < out.nRow; r++ {
		for c := 0; c < out.nCol; c++ {
			for i := 0; i < b.nRow; i++ {
				out.Set(r, c, out.Get(r, c)+(m.Get(r, i)*b.Get(i, c)))
			}
		}
	}

	return out
}

func (m *Matrix) Equals(data [][]int) bool {
	if m.nRow != len(data) {
		return false
	}

	if m.nCol != len(data[0]) {
		return false
	}

	for r, cols := range data {
		for c, v := range cols {
			if m.Get(r, c) != v {
				return false
			}
		}
	}

	return true
}
