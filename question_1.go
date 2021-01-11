package leetcode

func twoSum(nums []int, target int) []int {
    need := make(map[int]int)
    for idx, value := range nums {
         diff := target - value
         if v, ok := need[diff]; ok {
             result := []int{v, idx}
             return result
         } else {
             need[value] = idx
         }
    }
    return nil
}
