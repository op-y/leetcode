package leetcode

func searchInsert(nums []int, target int) int {
    for idx, value := range nums {
        if value >= target {
            return idx
        }
    }
    return len(nums)
}
