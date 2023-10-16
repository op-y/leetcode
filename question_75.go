package leetcode

func sortColors(nums []int)  {
	threeWayQuickSort(nums, 0, len(nums)-1)
}

func threeWayQuickSort(nums []int, lo, hi int) {
	if hi <= lo {
		return
	}
	v := nums[lo]
	lt := lo
	i := lo + 1
	gt := hi
	for i <= gt {
		if nums[i] < v {
			nums[lt], nums[i] = nums[i], nums[lt]
			lt++
			i++
		} else if nums[i] > v {
			nums[gt], nums[i] = nums[i], nums[gt]
			gt--
		} else {
			i++
		}
	}
	threeWayQuickSort(nums, lo, lt-1)
	threeWayQuickSort(nums, gt+1, hi)
}
