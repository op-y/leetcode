package leetcode

func missingNumber1836(nums []int) int {
	return search(nums, 0, len(nums)-1)
}

func search(a []int, lo, hi int) int {
	if lo > hi {
		return lo
	}
	mid := (lo + hi) / 2
	if a[mid] > mid {
		return search(a, lo, mid-1)
	} else {
		return search(a, mid+1, hi)
	}
}
