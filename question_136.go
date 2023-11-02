package leetcode

func singleNumber(nums []int) int {
    for idx, value := range nums {
        if idx == 0 {
            continue
        }
        nums[0] ^= value // 全部按位异或
    }
    return nums[0]
}
