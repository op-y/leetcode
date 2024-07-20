package leetcode

func incremovableSubarrayCount(nums []int) int {
    count := 0
    l, r := 1, len(nums)

    for l < r && nums[l] > nums[l-1] {
        l ++
    }

    count += l
    if l < r {
        count++
    }

    r = len(nums) - 2
    for r >= 0 {
        for l > 0 && nums[l-1] >= nums[r+1] {
            l--
        }
        count += l
        if l <= r {
            count++
        }
        if nums[r] >= nums[r+1] {
            break
        }
        r--
    }

    return count
}
