package leetcode

import (
	"math"
)

func numSquares(n int) int {
	f := make([]int , n+1)
	for i:=1; i<=n; i++ {
		minx := math.MaxInt32
		for j:=1; j*j <= i; j++ {
			minx = min(minx, f[i-j*j])
		}
		f[i] = minx + 1
	}
	return f[n]
}
