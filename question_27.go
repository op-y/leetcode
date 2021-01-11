package leetcode

func removeElement(nums []int, val int) int {
    if len(nums) == 0 {
        return 0
    }
    pos := 0
    for _, value := range nums {
        if value != val {
            nums[pos] = value
            pos++
        }
    }
    return pos
}
