package leetcode

func findMin154(nums []int) int {
    if len(nums) == 1 {
        return nums[0]
    }
    l, h := 0, len(nums) - 1
    for l < h {
        mid := (l + h) / 2
		if nums[mid] > nums[h] {
			l = mid + 1
		} else if nums[mid] < nums[h] {
			h = mid
        } else {
			h--
        }
    }
    return nums[l]
}
