package elm

type Matrix [][]int

func (m Matrix) Spiral() []int {
	n := len(m) * len(m)
	ch := make(chan int, n)

	m.spiral(0, ch)
	close(ch)
	var out []int
	for v := range ch {
		out = append(out, v)
	}

	return out
}

func (m Matrix) spiral(l int, ch chan int) {
	n := len(m)

	col := n - 1 - l
	if col == 1 {
		ch <- m[n/2][n/2]
		return
	}

	if col < 0 {
		return
	}

	for i := l; i < col; i++ {
		ch <- m[l][i]
	}

	for i := l; i < col; i++ {
		ch <- m[i][col]
	}

	for i := col; i > l; i-- {
		ch <- m[col][i]
	}

	for i := col; i > l; i-- {
		ch <- m[i][l]
	}

	m.spiral(l+1, ch)

}

// func TestSpiral(t *testing.T) {
// 	m := Matrix([][]int{
// 		[]int{1, 2, 3},
// 		[]int{4, 5, 6},
// 		[]int{7, 8, 9},
// 	})
// 	assert := require.New(t)
// 	assert.Equal(m.Spiral(), []int{1, 2, 3, 6, 9, 8, 7, 4, 5})

// 	m = Matrix([][]int{
// 		[]int{1, 2, 3, 4},
// 		[]int{5, 6, 7, 8},
// 		[]int{9, 10, 11, 12},
// 		[]int{13, 14, 15, 16},
// 	})
// 	// assert := require.New(t)
// 	assert.Equal(m.Spiral(), []int{1, 2, 3, 4, 8, 12, 17, 15, 14, 13, 9, 5, 6, 7, 11, 10})
// }
