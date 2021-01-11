package leetcode

func rob(nums []int) int {
    stat0 := 0
    stat1 := 0
    stat2 := 0
    if nums == nil {
        return 0
    }
    if len(nums) == 0 {
        return 0
    }
    stat0 = nums[0]
    if len(nums) == 1 {
        return stat0
    }
    if nums[0] > nums[1] {
        stat1 = nums[0]
    } else {
        stat1 = nums[1]
    }
    if len(nums) == 2 {
        return stat1
    }

    for i:=2; i<len(nums); i++ {
        if stat0 + nums[i] > stat1 {
            stat2 = stat0 + nums[i]
        } else {
            stat2 = stat1
        }
        stat0 = stat1
        stat1 = stat2
    }
    return stat2
}
