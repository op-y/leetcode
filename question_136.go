package leetcode

func singleNumber(nums []int) int {
    for idx, value := range nums {
        if idx == 0 {
            continue
        }
        nums[0] ^= value
    }
    return nums[0]
}
