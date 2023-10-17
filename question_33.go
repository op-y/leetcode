package leetcode

func search33(nums []int, target int) int {
    if nums == nil || len(nums) == 0 {
        return -1
    }
    l, h := 0, len(nums) - 1
    for l <= h {
        mid := (l + h) / 2
        if nums[mid] == target {
            return mid
        }
        if nums[0] <= nums[mid] {
            if nums[0] <= target && target < nums[mid] {
                h = mid - 1
            } else {
                l = mid + 1
            }
        } else {
            if nums[mid] < target && target <= nums[len(nums) - 1] {
                l = mid + 1
            } else {
                h = mid - 1
            }
        }
    }
    return -1
}
