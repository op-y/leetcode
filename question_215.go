package leetcode


func findKthLargest(nums []int, k int) int {
	n := len(nums)
	return quickselect(nums, 0, n - 1, n - k) // 因为是第k大，排序是升序，所以这里是n-k
}


// k 是目标下标
func quickselect(nums []int, l, r, k int) int {
	if l == r {
		return nums[k]
	}

	partition := nums[l]
	i, j := l - 1, r + 1
	for i < j { // 一轮分区过程
		for i++; nums[i]<partition; i++{}
		for j--; nums[j]>partition; j--{}
		if i < j {
			nums[i], nums[j] = nums[j], nums[i]
		}
	}
	if k <= j {
		return quickselect(nums, l, j, k) // k在左侧
	} else {
		return quickselect(nums, j + 1, r, k) // k在右侧
	}
}
