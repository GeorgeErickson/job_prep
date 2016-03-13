package elm

import (
	"testing"

	"gee.io/job_prep/ds"
)

func Pivot(a []int, idx int) []int {
	var less []int
	var equal []int
	var greater []int
	if idx >= len(a) || idx < 0 {
		return less
	}

	vIdx := a[idx]
	for _, v := range a {
		if v == vIdx {
			equal = append(equal, v)
		} else if v > vIdx {
			greater = append(greater, v)
		} else {
			less = append(less, v)
		}
	}

	less = append(less, equal...)
	less = append(less, greater...)
	return less
}

func PivotInPlace(a []int, idx int) []int {
	n := len(a)
	li, ei, gi := 0, 0, n-1
	vIdx := a[idx]

	for ei < gi+1 {
		v := a[ei]
		switch {
		case v == vIdx:
			ei++
		case v > vIdx:
			a[ei], a[gi] = a[gi], a[ei]
			gi--
		case v < vIdx:
			a[ei], a[li] = a[li], a[ei]
			li++
			ei++
		}
	}
	return a
}

func PivotOk(a []int, val int) bool {
	part := 0
	for _, v := range a {
		switch part {
		case 0:
			if v >= val {
				part++
			}
		case 1:
			// e.g. val=3 1,1,3,1
			if v < val {
				return false
			}
			if v > val {
				part++
			}
		case 2:
			// e.g. val=3 1,1,3,4,3
			if v <= val {
				return false
			}
		}
	}

	return true
}

func TestPivotOk(t *testing.T) {
	for _, bad := range [][]int{{1, 1, 3, 1}, {1, 2, 3, 4, 3}} {
		if PivotOk(bad, 3) {
			t.Errorf("PivotOk should be false for %v", bad)
		}
	}

	for _, good := range [][]int{{4, 1, 2, 6, 9, 8, 9}, {1, 2, 3, 4, 6}} {
		if !PivotOk(good, 6) {
			t.Errorf("PivotOk should be true for %v", good)
		}
	}
}

func TestPivot(t *testing.T) {
	// assert := require.New(t)
	a := ds.RandomIntList(100)
	i := 0

	v := a[i]
	act := PivotInPlace(a, i)
	if !PivotOk(act, v) {
		t.Errorf("Invalid pivot v: %v %v", v, act)
	}

}
