package arrays

func BinarySearch(in []int, val int) int {
	l := 0
	h := len(in)

	for l <= h {
		m := l + ((h - l) / 2)

		v := in[m]

		switch {
		case v == val:
			return m
		case v < val:
			l = m
		case v > val:
			h = m
		}
	}

	return -1
}
