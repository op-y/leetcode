package leetcode

func searchRange(nums []int, target int) []int {
	l := len(nums)
	if l == 0 {
		return []int{-1, -1}
	}
	if l == 1 && target == nums[0] {
		return []int{0, 0}
	}

	idx := search34(nums, target, 0, l-1)
	if idx == -1 {
		return []int{-1, -1}
	} else {
		first := search34left(nums, target, 0, idx)
		last := search34right(nums, target, idx, l-1)
		return []int{first, last}
	}
}

func search34(nums []int, target, lo, hi int) int {
	if lo > hi {
		return -1
	}
	mid := (lo + hi) / 2
	if nums[mid] == target {
		return mid
	} else if  nums[mid] > target {
		return search34(nums, target, lo, mid-1)
	} else {
		return search34(nums, target, mid+1, hi)
	}
}

func search34left(nums []int, target, lo, hi int) int {
	if lo > hi {
		return -1
	}
	mid := (lo + hi) / 2
	if nums[mid] == target {
		if mid > lo && nums[mid] > nums[mid-1] {
			return mid
		}
		if mid == lo {
			return mid
		}
		return search34left(nums, target, lo, mid-1)
	} else if nums[mid] < target {
		return search34left(nums, target, mid+1, hi)
	} else {
		return search34left(nums, target, lo, mid-1)
	}
}

func search34right(nums []int, target, lo, hi int) int {
	if lo > hi {
		return -1
	}
	mid := (lo + hi) / 2
	if nums[mid] == target {
		if mid < hi && nums[mid] < nums[mid+1] {
			return mid
		}
		if mid == hi {
			return mid
		}
		return search34right(nums, target, mid+1, hi)
	} else if nums[mid] > target  {
		return search34right(nums, target, lo, mid-1 )
	} else {
		return search34right(nums, target, mid+1, hi)
	}
}
