package cracking

import "gee.io/job_prep/ds"

func powerBySquaring(x, n int) int {
	out := x
	tail := 0
	for n > 1 {
		if n&1 == 1 {
			tail++
			n--
		} else {
			out *= out
			n = n / 2
		}
	}

	for i := tail; i > 0; i-- {
		out = out * x
	}

	return out
}

// 0, 1, 0     0, 1, 0   00 * 00 + 01 * 10
// 0, 0, 1  *  0, 0, 1 =
// 1, 1, 1     1, 1, 1
func square33(m [3][3]int) [3][3]int {
	out := [3][3]int{
		{0, 0, 0},
		{0, 0, 0},
		{0, 0, 0},
	}

	for r := 0; r < 3; r++ {
		for c := 0; c < 3; c++ {
			for i := 0; i < 3; i++ {
				out[r][c] += (m[i][c] * m[r][i])
			}
		}
	}

	return out
}

func square22(left [2][2]int, right [2][2]int) [2][2]int {
	out := [2][2]int{
		{0, 0},
		{0, 0},
	}

	for r := 0; r < 2; r++ {
		for c := 0; c < 2; c++ {
			for i := 0; i < 2; i++ {
				out[r][c] += (left[i][c] * right[r][i])
			}
		}
	}

	return out
	// return [3][3]int{
	// 	{
	// 		(m[0][0] * m[0][0]) + (m[0][1] * m[1][0]) + (m[0][2] * m[2][0]),
	// 		(m[0][0] * m[0][1]) + (m[0][1] * m[1][1]) + (m[0][2] * m[2][1]),
	// 		(m[0][0] * m[0][2]) + (m[0][2] * m[1][2]) + (m[0][2] * m[2][2]),
	// 	},
	// 	{0, 0, 0},
	// 	{0, 0, 0},
	// }

	//1010
}

// 11 1
// 10 2
// 5 1
// 4 2
// 2

func fibBad(n int) int {
	if n == 0 || n == 1 {
		return n
	}

	return fibBad(n-1) + fibBad(n-2)
}

func multFib(a, b, ta, tb int) (int, int) {
	return ta*a + tb*b, ta*b + tb*(a+b)
}

func fib2(n int) int {
	if n == 0 || n == 1 {
		return n
	}

	a, b := 0, 1
	ta, tb := 1, 0

	for n > 1 {
		if n&1 == 1 {
			ta, tb = multFib(a, b, ta, tb)
			n--
		} else {
			a, b = multFib(a, b, a, b)
			n = n / 2
		}
	}
	_, out := multFib(a, b, ta, tb)
	return out
}

func fib(n int) int {
	if n == 0 || n == 1 {
		return n
	}

	m := ds.MatrixFrom([][]int{
		{0, 1},
		{1, 1},
	})

	tail := ds.MatrixFrom([][]int{
		{1, 0},
		{0, 1},
	})

	for n > 1 {
		if n&1 == 1 {
			tail = tail.Product(m)
			n--
		} else {
			m = m.Square()
			n /= 2
		}
	}

	return m.Product(tail).Get(0, 1)
}

// func q81(n int) {
// 	m := [3][3]int{
// 		{0, 1, 0},
// 		{0, 0, 1},
// 		{1, 1, 1},
// 	}

// }
