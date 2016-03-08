package ds

import "math/rand"

func randomIntList(n int) []int {
	out := make([]int, n)
	for i := 0; i < n; i++ {
		out[i] = rand.Int()
	}

	return out
}
