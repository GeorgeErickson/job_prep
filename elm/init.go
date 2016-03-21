package elm

import (
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
}

func NChooseK(n, k int) int {
	if k > n/2 {
		k = n - k
	}

	num := 1
	d := 1
	for i := 0; i < k; i++ {
		num *= (n - i)
		d *= (i + 1)
	}
	return num / d
}
