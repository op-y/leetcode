package leetcode

import (
	"math"
)

func findPeakElement(nums []int) int {
	l := len(nums)

    get := func(i int) int {
        if i == -1 || i == l {
            return math.MinInt64
        }
        return nums[i]
    }

	
	left, right := 0, l-1
	for {
		mid := (left + right) / 2
		if get(mid-1) < get(mid) && get(mid) > get(mid+1) {
			return mid
		}
		if get(mid) < get(mid+1) {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
}

