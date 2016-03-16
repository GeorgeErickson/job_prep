package elm

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func MaxSingleProfitBrute(a []int) int {
	var cMax int
	for i, buy := range a {
		for _, sell := range a[i:] {
			v := sell - buy
			if v > cMax {
				cMax = v
			}
		}
	}

	return cMax
}

func MaxSingleProfit(a []int) int {
	if len(a) == 0 {
		return 0
	}
	cMin := a[0]
	cMax := 0
	for _, v := range a {
		max := v - cMin
		if max > cMax {
			cMax = max
		}
		if v < cMin {
			cMin = v
		}
	}
	return cMax
}

func MaxDoubleProfit(a []int) int {
	firstSaleMax := make([]int, len(a))

	// calc max profit for sale on day i,
	minPriceSeen := a[0]
	maxProfitSeen := 0
	for i, iPrice := range a {
		iProfit := iPrice - minPriceSeen

		if iProfit > maxProfitSeen {
			maxProfitSeen = iProfit
		}

		if iPrice < minPriceSeen {
			minPriceSeen = iPrice
		}
		firstSaleMax[i] = maxProfitSeen
	}

	// max profit if buy after day i
	maxPriceSeen := a[len(a)-1]
	maxProfitSeen = 0
	for i := len(a) - 1; i > 0; i-- {
		iPrice := a[i]
		if iPrice > maxPriceSeen {
			maxPriceSeen = iPrice
		}

		iProfit := maxPriceSeen - iPrice + firstSaleMax[i]
		if iProfit > maxProfitSeen {
			maxProfitSeen = iProfit
		}
	}

	return maxProfitSeen
}

func TestMaxSingleProfit(t *testing.T) {
	assert := require.New(t)
	a := []int{310, 315, 275, 260, 270, 290, 230, 255, 250}
	assert.Equal(30, MaxSingleProfitBrute(a))
	assert.Equal(30, MaxSingleProfit(a))
}

func TestMaxDoubleProfit(t *testing.T) {
	assert := require.New(t)
	a := []int{310, 315, 275, 260, 270, 290, 230, 255, 250}
	assert.Equal(55, MaxDoubleProfit(a))

	// sort.IntList
	// assert.Equal(30, MaxSingleProfit(a))
}
