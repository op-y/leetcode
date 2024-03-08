package leetcode

import (
	"math"
)

func coinChange(coins []int, amount int) int {
	f := make([]int, amount+1)
	l := len(coins)

	for i := 1; i <= amount; i++ {
		minx := math.MaxInt32
		found := false

		for j := 0; j < l; j++ {
			if i-coins[j] >= 0 && f[i-coins[j]] != -1 {
				minx = min(minx, f[i-coins[j]])
				found = true
			}
		}
		if found {
			f[i] = minx + 1
		} else {
			f[i] = -1
		}
	}

	return f[amount]
}
