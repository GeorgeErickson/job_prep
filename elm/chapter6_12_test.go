package elm

import "math/rand"

//
func SampleOfflineData(a []int, count int) []int {
	if count >= len(a) {
		return a
	}

	for k := 0; k <= count; k++ {
		i := rand.Intn(len(a)-k) + k
		a[i], a[k] = a[k], a[i]
	}

	return a[:count+1]
}
